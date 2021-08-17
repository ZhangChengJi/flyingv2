package etcd

import (
	"flyingv2/logs"
	clientv3 "go.etcd.io/etcd/client/v3"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"time"
)

const (
	defaultEndpoints = "http://192.168.1.227:2379"
	// The short keepalive timeout and interval have been chosen to aggressively
	// detect a failed etcd server without introducing much overhead.
	keepaliveTime    = 30 * time.Second
	keepaliveTimeout = 10 * time.Second

	// dialTimeout is the timeout for failing to establish a connection.
	// It is set to 20 seconds as times shorter than that will cause TLS connections to fail
	// on heavily loaded arm64 CPUs (issue #64649)
	dialTimeout = 20 * time.Second
)

func NewEtcdClient() (*clientv3.Client, error) {
	var endpoints string
	//flag.StringVar(&endpoints, "ETCD_ENDPOINT", "", "etcd 连接")
	//flag.Parse()
	if endpoints == "" {
		logs.L.Warn("Use default etcd connection: " + defaultEndpoints)
		endpoints = defaultEndpoints
	}

	cfg := clientv3.Config{
		Endpoints:            []string{endpoints},
		DialTimeout:          dialTimeout,
		DialKeepAliveTime:    keepaliveTime,
		DialKeepAliveTimeout: keepaliveTimeout,
		DialOptions: []grpc.DialOption{
			grpc.WithBlock(), // block until the underlying connection is up
		},
	}
	e, err := clientv3.New(cfg)
	if err != nil {
		logs.L.Error("etcd create connection failed:", zap.Error(err))
	}
	return e, err
}
