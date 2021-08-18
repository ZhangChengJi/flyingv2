package factory

import (
	"flyingv2/core"
	"flyingv2/core/etcd"
)

func Create(prefix string) core.Interface {
	return newETCD3Storage(prefix)
}

//初始化了etcd
func newETCD3Storage(prefix string) core.Interface {
	client, err := etcd.NewEtcdClient()
	if err != nil {
		return nil
	}
	return etcd.New(client, prefix)
}
