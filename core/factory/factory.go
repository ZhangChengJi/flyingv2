package factory

import (
	"flyingv2/core"
	"flyingv2/core/etcd"
)

func Create() (core.Interface, error) {
	return newETCD3Storage()
}

//初始化了etcd
func newETCD3Storage() (core.Interface, error) {
	client, err := etcd.NewEtcdClient()
	if err != nil {
		return nil, err
	}
	return etcd.New(client, "/"), nil
}
