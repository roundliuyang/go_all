package discover

import (
	"context"
	clientv3 "go.etcd.io/etcd/client/v3"
	"google.golang.org/grpc/resolver"
	"log"
	"time"
)

/*
客户端首先需要实现接口resolver.Resolver，其中方法Build()用于创建一个etcd解析器，grpc.Dial()会同步调用该方法，
解析器需要根据key前缀监听etcd中服务地址列表的变化并更新本地列表watch。然后注册解析器，创建gRPC句柄，使用轮询负载均衡请求服务。
*/
type etcdResolverBuilder struct {
	etcdClient *clientv3.Client
}

func NewEtcdResolverBuilder() *etcdResolverBuilder {

	// 创建etcd客户端连接
	etcdClient, err := clientv3.New(clientv3.Config{
		Endpoints:   []string{"127.0.0.1:2379"},
		DialTimeout: 5 * time.Second,
	})

	if err != nil {
		log.Println("client get etcd failed,error", err)
		panic(err)
	}

	return &etcdResolverBuilder{
		etcdClient: etcdClient,
	}
}

func (erb *etcdResolverBuilder) Build(target resolver.Target, cc resolver.ClientConn,
	opts resolver.BuildOptions) (resolver.Resolver, error) {

	// 获取指定前缀的etcd节点值
	// /ns->/ns/order-service-1   /ns/order-service-2
	prefix := "/" + target.URL.Scheme

	log.Println(prefix)

	// 获取 etcd 中服务保存的ip列表
	res, err := erb.etcdClient.Get(context.Background(), prefix, clientv3.WithPrefix())

	if err != nil {
		log.Println("Build etcd get addr failed; err:", err)
		return nil, err
	}

	ctx, cancelFunc := context.WithCancel(context.Background())

	es := &etcdResolver{
		cc:         cc,
		etcdClient: erb.etcdClient,
		ctx:        ctx,
		cancel:     cancelFunc,
		scheme:     target.URL.Scheme,
	}

	// 将获取到的ip和port保存到本地的map中
	log.Printf("etcd res:%+v\n", res)
	for _, kv := range res.Kvs {
		es.store(kv.Key, kv.Value)
	}

	// 更新拨号里的ip列表
	es.updateState()

	// 监听etcd中的服务是否变化

	go es.watcher()
	return es, nil
}

func (erb *etcdResolverBuilder) Scheme() string {
	return "etcd"
}
