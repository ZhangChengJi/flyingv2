package etcd

import (
	"context"
	"errors"
	"flyingv2/internal/core"
	"flyingv2/internal/core/model"
	"flyingv2/logs"
	"fmt"
	"go.etcd.io/etcd/client/v3"
	"go.uber.org/zap"
	"path"
	"time"
)

type Store struct {
	Client *clientv3.Client
	Prefix string
}

const (
	priority    = 10  // default priority when nothing is set
	ttl         = 300 // default ttl when nothing is set
	etcdTimeout = 5 * time.Second
)

func New(prefix string) core.Interface {
	return newStorage(Client, prefix)
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
	kv := clientv3.NewKV(s.Client)
	txn := kv.Txn(ctx)
	txn.If(clientv3.Compare(clientv3.CreateRevision(key), "=", 0)).
		Then(clientv3.OpPut(key, value, clientv3.WithPrevKV())).
		Else(clientv3.OpGet(key))
	txnResp, err := txn.Commit()
	if err != nil {
		fmt.Println(err)
		return nil // 没有问题
	}
	// 判断是否抢到了锁
	if !txnResp.Succeeded {
		fmt.Println("锁被占用:", string(
			txnResp.Responses[0].GetResponseRange().Kvs[0].Value))
		return nil
	}

	return nil
}
func (s *Store) Get(ctx context.Context, key string) (interface{}, error) {
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

	return string(r.Kvs[0].Value), nil
}

func (s *Store) List(ctx context.Context, opts *model.ListOptions) (list *model.PageList, err error) {

	options := make([]clientv3.OpOption, 0, 4)
	key := opts.Key
	options = append(options, clientv3.WithLimit(opts.PageInfo.PageSize))
	options = append(options, clientv3.WithSort(clientv3.SortByKey, clientv3.SortAscend))
	if key != "" { //特殊指定key查询
		key = path.Join(s.Prefix, key)
		options = append(options, clientv3.WithPrefix())
	} else { //范围查询,使用默认前缀
		key = s.Prefix
	}

	//TODO: 坑
	options = append(options, clientv3.WithRange(clientv3.GetPrefixRangeEnd(key)))
	var i int64
	var total int64

	for {
		resp, _ := s.Client.Get(ctx, key, options...)
		i++
		if i == 1 {
			total = resp.Count
		}
		if len(resp.Kvs) == 0 {
			break
		}
		if i >= opts.Page || !resp.More {
			list.Unmarshal(total, resp)
			break
		}
		key = string(resp.Kvs[len(resp.Kvs)-1].Key) + "\x00"
	}

	return list, err

}

func (s *Store) Update(ctx context.Context, key string, value string) error {
	_, err := s.Client.Put(ctx, key, value)
	if err != nil {
		logs.L.Info("etcd put failed: %s", zap.Error(err))
		return fmt.Errorf("etcd put failed: %s", err)
	}
	return nil
}
