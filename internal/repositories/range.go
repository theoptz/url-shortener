package repositories

import (
	"context"
	"encoding/json"
	"errors"
	"strconv"

	"github.com/sirupsen/logrus"
	"github.com/theoptz/url-shortener/internal/utils"
	clientv3 "go.etcd.io/etcd/client/v3"
	"go.etcd.io/etcd/client/v3/concurrency"
)

const prefix = "/nodes/"
const nextValPrefix = "index"
const rangeLength = 100000
const startIndex = 62 * 62 * 62 * 62

type RangeRepo struct {
	client *clientv3.Client
	name   string
}

func (r *RangeRepo) getNode() string {
	return prefix + r.name
}

func (r *RangeRepo) Get(ctx context.Context) (*RangeItem, error) {
	val, err := r.client.Get(ctx, r.getNode())
	if err != nil {
		return nil, err
	}

	var item RangeItem

	if val.Count > 0 {
		err = json.Unmarshal(val.Kvs[0].Value, &item)
		if err != nil {
			return nil, err
		}

		logrus.Println("value exists", item)

		return &item, nil
	}

	return r.GetNext(ctx)
}

func (r *RangeRepo) GetNext(parentContext context.Context) (*RangeItem, error) {
	ctx, cancel := context.WithCancel(parentContext)
	defer cancel()

	session, err := concurrency.NewSession(r.client)
	if err != nil {
		return nil, err
	}
	defer session.Close()

	mu := concurrency.NewMutex(session, nextValPrefix)

	err = mu.Lock(ctx)
	if err != nil {
		return nil, err
	}

	index, err := r.getIndex(ctx)
	if err != nil {
		return nil, err
	}

	nextIndex := index + rangeLength
	val := RangeItem{
		Start: index,
		End:   nextIndex - 1,
	}

	err = r.setIndexAndNode(ctx, nextIndex, &val)
	if err != nil {
		return nil, err
	}

	err = mu.Unlock(ctx)
	if err != nil {
		return nil, err
	}

	return &val, nil
}

func (r *RangeRepo) getIndex(ctx context.Context) (int64, error) {
	nextVal, err := r.client.Get(ctx, nextValPrefix)
	if err != nil {
		return 0, err
	}

	var res int64
	if nextVal.Count > 0 {
		err = json.Unmarshal(nextVal.Kvs[0].Value, &res)
		if err != nil {
			return 0, err
		} else if res > startIndex {
			return res, nil
		}
	}

	return startIndex, nil
}

func (r *RangeRepo) setIndexAndNode(ctx context.Context, val int64, rangeItem *RangeItem) error {
	if val < startIndex {
		return errors.New("incorrect value for next index")
	}

	data, err := json.Marshal(rangeItem)
	if err != nil {
		return err
	}

	tnx := r.client.Txn(ctx)

	_, err = tnx.Then(
		clientv3.OpPut(nextValPrefix, strconv.FormatInt(val, 10)),
		clientv3.OpPut(r.getNode(), utils.GetString(data)),
	).Commit()

	return err
}

func NewRangRepo(client *clientv3.Client, name string) *RangeRepo {
	return &RangeRepo{
		client: client,
		name:   name,
	}
}
