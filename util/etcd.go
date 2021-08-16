package util
import(
	"flag"
	"fmt"
	clientv3 "go.etcd.io/etcd/client/v3"
	"google.golang.org/grpc"
	"time"
)

type etcd struct{
	cli *clientv3.Client
}

func(e etcd) run(){

}

func newEtcdClient() *etcd{
     var endpoints string
     flag.StringVar(&endpoints,"ETCD_ENDPOINT","","etcd 连接")
     flag.Parse()
     if endpoints==""{
     	fmt.Println("Use default etcd connection")
     	endpoints="127.0.0.1:2379"
	 }
	cli,err:=clientv3.New(
		clientv3.Config{
			Endpoints: [] string{endpoints},
			DialTimeout: 2 * time.Second,
			DialOptions: []grpc.DialOption{
				grpc.WithBlock(), // block until the underlying connection is up
			},
		})
	 if err != nil {
	 	fm
	 }
	e:=&etcd{
		cli:
	}
	return e

}
