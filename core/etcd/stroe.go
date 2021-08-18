package etcd

import (
	"context"
	"errors"
	"flyingv2/core"
	"fmt"
	clientv3 "go.etcd.io/etcd/client/v3"
	"path"
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

func (s *Store) Set(ctx context.Context, key string, value string) error {
	ctx, cancel := context.WithTimeout(ctx, etcdTimeout)
	defer cancel()
	key = path.Join(s.Prefix, key)
	aas, _ := s.Client.Put(ctx, key, value, clientv3.WithPrevKV())
	fmt.Println(aas)
	return nil
}
func (s *Store) Get(ctx context.Context, key string) ([]byte, error) {
	ctx, cancel := context.WithTimeout(ctx, etcdTimeout)
	defer cancel()
	//if recursive {
	key = path.Join(s.Prefix, key)
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

	r, err := s.Client.Get(ctx, key)
	if err != nil {
		return []byte{}, err
	}
	if r.Count == 0 {
		return []byte{}, errKeyNotFound
	}

	return r.Kvs[0].Value, nil
}

func (s *Store) List(ctx context.Context, key string) ([][]byte, error) {

	//if !strings.HasSuffix(key, "/") {
	//	key += "/"
	//}
	options := make([]clientv3.OpOption, 0, 4)
	options = append(options, clientv3.WithLimit(5))
	if key == "" {
		options = append(options, clientv3.WithPrefix())
	} else {
		options = append(options, clientv3.WithFromKey())
	}
	key = path.Join(s.Prefix, key)
	//keyPrefix := key
	rangeEnd := clientv3.GetPrefixRangeEnd(keyPrefix)
	options = append(options, clientv3.WithRange(rangeEnd))
	options = append(options, clientv3.WithSort(clientv3.SortByKey, clientv3.SortAscend))

	//options = append(options, clientv3.WithPrefix())
	var lastKey []byte
	for {
		getResp := new(clientv3.GetResponse)
		//key = key + "\x00"
		getResp, _ = s.Client.Get(ctx, key, options...)

		for _, kv := range getResp.Kvs {
			fmt.Println(kv)
			lastKey = kv.Key
		}
		key = string(lastKey) + "\x00"
		break
	}

	//get, err := e.Client.Get(ctx, path, clientv3.WithPrefix())
	//if err != nil {
	//	return nil, err
	//}
	//fmt.Println(get)
	return [][]byte{}, nil

}
