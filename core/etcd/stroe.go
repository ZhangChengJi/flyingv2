package etcd

import (
	"context"
	"errors"
	clientv3 "go.etcd.io/etcd/client/v3"
	"strings"
	"time"
)

type Stroe struct {
	Client *clientv3.Client
}

const (
	priority    = 10  // default priority when nothing is set
	ttl         = 300 // default ttl when nothing is set
	etcdTimeout = 5 * time.Second
)

var errKeyNotFound = errors.New("key not found")

func (e *Stroe) Get(ctx context.Context, path string, recursive bool) (*clientv3.GetResponse, error) {
	ctx, cancel := context.WithTimeout(ctx, etcdTimeout)
	defer cancel()
	if recursive {
		if !strings.HasSuffix(path, "/") {
			path = path + "/"
		}
		r, err := e.Client.Get(ctx, path, clientv3.WithPrefix())
		if err != nil {
			return nil, err
		}
		if r.Count == 0 {
			path = strings.TrimSuffix(path, "/")
			r, err = e.Client.Get(ctx, path)
			if err != nil {
				return nil, err
			}
			if r.Count == 0 {
				return nil, errKeyNotFound
			}
		}
		return r, nil
	}

	r, err := e.Client.Get(ctx, path)
	if err != nil {
		return nil, err
	}
	if r.Count == 0 {
		return nil, errKeyNotFound
	}
	return r, nil
}
