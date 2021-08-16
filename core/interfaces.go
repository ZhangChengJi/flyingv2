package core

import (
	"context"
	clientv3 "go.etcd.io/etcd/client/v3"
)

type EtcdClient interface {
	Get(ctx context.Context, path string, recursive bool) (*clientv3.GetResponse, error)
}
