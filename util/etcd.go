package util
import(
	clientv3 "go.etcd.io/etcd/client/v3"
)

type etcd struct{
	cli *clientv3.Client
}

func run(e etcd){

}

func newEtcdClient() *etcd{

	e:=&etcd{
		cli: clientv3.New(
			clientv3.Config{
				Endpoints:

		})
	}
	return e

}
