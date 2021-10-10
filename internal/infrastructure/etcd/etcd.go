package etcd

import (
	"sync"
	"time"

	clientv3 "go.etcd.io/etcd/client/v3"
)

var client *clientv3.Client

func GetClient() *clientv3.Client {
	return client
}

func Connect(addr string) (*clientv3.Client, error) {
	var err error
	once := sync.Once{}

	once.Do(func() {
		client, err = clientv3.New(clientv3.Config{
			Endpoints:   []string{addr},
			DialTimeout: 5 * time.Second,
		})
	})

	return client, err
}
