package counter

import (
	"context"
	"errors"
	"sync"
	"sync/atomic"
	"time"

	"github.com/theoptz/url-shortener/internal/interfaces/irepositories"
	"github.com/theoptz/url-shortener/internal/repositories"
)

const max = 62 * 62 * 62 * 62 * 62 * 62 * 62

type Counter struct {
	rangeRepo irepositories.RangeRepository

	rangeItem repositories.RangeItem
	index     int64

	mu sync.Mutex
}

func (c *Counter) updateRange() error {
	c.mu.Lock()
	defer c.mu.Unlock()

	if atomic.LoadInt64(&c.index) >= c.rangeItem.End {
		ctx, cancel := context.WithTimeout(context.TODO(), time.Second*2)
		defer cancel()

		rangeItem, err := c.rangeRepo.GetNext(ctx)
		if err != nil {
			return err
		}

		c.rangeItem = *rangeItem
		atomic.StoreInt64(&c.index, c.rangeItem.Start-1)
	}

	return nil
}

func (c *Counter) Inc() (int64, error) {
	v := atomic.AddInt64(&c.index, 1)
	end := atomic.LoadInt64(&c.rangeItem.End)

	if v > max {
		return 0, errors.New("limit exceeded")
	} else if v > end {
		err := c.updateRange()
		if err != nil {
			return 0, err
		}

		return c.Inc()
	}

	return v, nil
}

func NewCounter(rangeRepo irepositories.RangeRepository, index int64, rangeItem repositories.RangeItem) *Counter {
	if index == 0 && rangeItem.Start > 0 {
		index = rangeItem.Start - 1
	}

	return &Counter{
		rangeRepo: rangeRepo,
		index:     index,
		rangeItem: rangeItem,
	}
}
