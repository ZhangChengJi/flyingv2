package etcd

import (
	"context"
	"errors"
	"flyingv2/core"
	"fmt"
	clientv3 "go.etcd.io/etcd/client/v3"
	"time"
)

var AA *Store

type Store struct {
	Client *clientv3.Client
	Prefix string
}

const (
	priority    = 10  // default priority when nothing is set
	ttl         = 300 // default ttl when nothing is set
	etcdTimeout = 5 * time.Second
)

var Storage core.Interface

func New(client *clientv3.Client, prefix string) core.Interface {
	return newStorage(client, prefix)
}
func newStorage(client *clientv3.Client, prefix string) *Store {
	return &Store{
		Client: client,
		Prefix: prefix,
	}
}

var errKeyNotFound = errors.New("key not found")

func (e *Store) Set(ctx context.Context, path string, value string) error {
	ctx, cancel := context.WithTimeout(ctx, etcdTimeout)
	defer cancel()
	aas, _ := e.Client.Put(ctx, path, value, clientv3.WithPrevKV())
	fmt.Println(aas)
	return nil
}
func (e *Store) Get(ctx context.Context, path string) ([]byte, error) {
	ctx, cancel := context.WithTimeout(ctx, etcdTimeout)
	defer cancel()
	//if recursive {
	//	if !strings.HasSuffix(path, "/") {
	//		path = path + "/"
	//	}
	//	r, err := e.Client.Get(ctx, path, clientv3.WithPrefix())
	//	if err != nil {
	//		return "", err
	//	}
	//	if r.Count == 0 {
	//		path = strings.TrimSuffix(path, "/")
	//		r, err = e.Client.Get(ctx, path)
	//		if err != nil {
	//			return "", err
	//		}
	//		if r.Count == 0 {
	//			return "", errKeyNotFound
	//		}
	//	}
	//	return string(r.Kvs[0].Value), nil
	//}

	r, err := e.Client.Get(ctx, path)
	if err != nil {
		return []byte{}, err
	}
	if r.Count == 0 {
		return []byte{}, errKeyNotFound
	}

	return r.Kvs[0].Value, nil
}

func (e *Store) GetAll(ctx context.Context, path string) ([][]byte, error) {
	//options := make([]clientv3.OpOption, 0, 4)
	//if !strings.HasSuffix(path, "/") {
	//	path += "/"
	//}
	get, err := e.Client.Get(ctx, path, clientv3.WithPrefix())
	if err != nil {
		return nil, err
	}
	fmt.Println(get)
	return [][]byte{}, nil

}
