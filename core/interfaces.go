package core

import (
	"context"
	"fmt"
	clientv3 "go.etcd.io/etcd/client/v3"
)

type Interface interface {
	Set(ctx context.Context, key string, value string) error
	Get(ctx context.Context, key string) ([]byte, error)
	List(ctx context.Context, key string, ops *ListOptions, list *PageList) error
}

type ListOptions struct {
	PageInfo
}

type PageInfo struct {
	Page     int64  `json:"page" form:"page"`
	PageSize int64  `json:"pageSize" form:"pageSize"`
	Key      string `json:"key" form:"key"`
}

type Page struct {
	Total    int64 `json:"total"`
	Page     int64 `json:"page"`
	PageSize int64 `json:"pageSize"`
}
type App struct {
	Name  string
	AppId string
}
type PageList struct {
	List interface{} `json:"list"`
	Page
}

func (p *PageList) Unmarshal(total int64, response *clientv3.GetResponse) {
	if len(response.Kvs) > 0 {
		list := make([]interface{}, 0)
		for _, kv := range response.Kvs {
			list = append(list, fmt.Sprintf("%v------%v", string(kv.Key), string(kv.Value)))
		}
		p.List = &list
		p.Total = total
		//p.PageSize

	}
	//response.Kvs
	//json.Unmarshal()
}
