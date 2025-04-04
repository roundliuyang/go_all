package main

import (
	"context"
	"fmt"
	clientv3 "go.etcd.io/etcd/client/v3"
	"log"
	"time"
)

/*
通过Grant，我们可以创建一个带有过期的时间的租约，但是，有些时候我们可能需要续租。通俗一点可以理解为：只要这个租约和ETCD保持(连接)KeepAlive，
那么就不会过期，如果断开，则会在5秒后过期。又或者可以理解为心跳机制。
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

	// 创建租约，TTL 为 5 秒
	leaseGrant, err := client.Grant(ctx, 5)
	if err != nil {
		log.Printf("grant error %v", err)
		return
	}

	// 将 key 写入 etcd，绑定这个租约
	_, err = client.Put(ctx, key, value, clientv3.WithLease(leaseGrant.ID))
	if err != nil {
		log.Printf("put error %v", err)
		return
	}

	// 启动自动续租，etcd 会周期性续这个租约
	keepaliveResponseChan, err := client.KeepAlive(ctx, leaseGrant.ID)
	if err != nil {
		log.Printf("KeepAlive error %v", err)
		return
	}

	for {
		ka := <-keepaliveResponseChan
		fmt.Println("ttl:", ka.TTL)
		//ttl: 5
		//ttl: 5
		//ttl: 5
	}
}
