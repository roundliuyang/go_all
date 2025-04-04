package main

import (
	"context"
	"fmt"
	clientv3 "go.etcd.io/etcd/client/v3"
	"log"
	"time"
)

/*
除了常规的操作外，我们还可以设置KEY-VALUE的过期时间，这里引入一个概念叫租约，将设置过期时间这一过程称为租约，租约期限到期则KEY-VALUE将删除
下面的代码表示：

	-创建一个租约，有效期为5秒
	-使用这个租约PUT一个新的KEY-VALUE
	-使用watcher监听这个KEY
	-5秒后将监听到DELETE事件，租约过期
*/
func main() {

	client, err := clientv3.New(clientv3.Config{
		Endpoints:   []string{"127.0.0.1:2379"},
		DialTimeout: 5 * time.Second,
	})
	if err != nil {
		log.Fatalln(err)
	}

	key := "/ns/service"
	value := "127.0.0.1:800"
	ctx := context.Background()

	// 获取一个租约 有效期为5秒
	leaseGrant, err := client.Grant(ctx, 5)
	if err != nil {
		log.Printf("put error %v", err)
		return
	}

	// PUT 租约期限为5秒
	_, err = client.Put(ctx, key, value, clientv3.WithLease(leaseGrant.ID))
	if err != nil {
		log.Printf("put error %v", err)
		return
	}

	// 监听变化 5秒后将监听到DELETE事件
	watcher(client, key)

}

func watcher(client *clientv3.Client, key string) {

	// 监听这个chan
	watchChan := client.Watch(context.Background(), key)

	for watchResponse := range watchChan {
		for _, event := range watchResponse.Events {
			fmt.Printf("Type:%s,Key:%s,Value:%s\n", event.Type, event.Kv.Key, event.Kv.Value)
			// Type:DELETE,Key:/ns/service,Value:
		}
	}

}
