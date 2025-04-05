package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/balancer/roundrobin"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/resolver"
	"io"
	"sgrpc/etcd/discover"
	"sgrpc/grpc/lb"
	"sgrpc/grpc/proto/hello"
	"strconv"
	"time"
)

func main() {
	// 注册自定义ETCD解析器
	etcdResolverBuilder := discover.NewEtcdResolverBuilder()
	resolver.Register(etcdResolverBuilder)

	// 使用自带的DNS解析器和负载均衡实现方式
	conn, err := grpc.Dial("etcd:///",
		grpc.WithBalancerName(roundrobin.Name),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithBalancerName(lb.WEIGHT_LOAD_BALANCE))
	if err != nil {
		panic(err)
	}

	helloServiceClient := hello.NewHelloServiceClient(conn)

	for {
		time.Sleep(2 * time.Second)
		err = SayHello(helloServiceClient) // 一元
		if err != nil {
			fmt.Println(err)
		}
	}
	/**
	2021/12/28 17:59:31 /etcd
	2021/12/28 17:59:31 etcd res:&{Header:cluster_id:11588568905070377092 member_id:128088275939295631 revision:254 raft_term:7  Kvs:[key:"/etcd/order-service-1" create_revision:254 mod_revision:254 version:1 value:"127.0.0.1:8001" lease:112447026244685706 ] More:false Count:1 XXX_NoUnkeyedLiteral:{} XXX_unrecognized:[] XXX_sizecache:0}
	2021/12/28 17:59:31 conn.UpdateState key[/etcd/order-service-1];val[127.0.0.1:8001]
	code:200 msg:"Hello! 8001一元RPC"
	code:200 msg:"Hello! 8001一元RPC"
	code:200 msg:"Hello! 8001一元RPC"
	*/

}

// 一元

func SayHello(client hello.HelloServiceClient) error {

	request := &hello.HelloRequest{
		Name: "一元RPC",
		Age:  18,
	}

	// 10秒超时
	ctx, cancelFunc := context.WithDeadline(context.Background(), time.Now().Add(10*time.Second))
	defer cancelFunc()

	helloResponse, err := client.SayHello(ctx, request)
	if err != nil {
		return err
	}

	fmt.Println(helloResponse)

	return nil
}

// 服务端流 服务端一直在发，所以客户端可以一直接收

func LotsOfReplies(client hello.HelloServiceClient) error {

	// 3秒超时
	ctx, cancelFunc := context.WithDeadline(context.Background(), time.Now().Add(20*time.Second))
	defer cancelFunc()

	repliesClient, err := client.LotsOfReplies(ctx, &hello.HelloRequest{
		Name: "服务端流式RPC",
		Age:  18,
	})

	if err != nil {
		return err
	}

	defer repliesClient.CloseSend()

	for {
		helloResponse, err := repliesClient.Recv()

		if err == io.EOF {
			return nil
		}

		if err != nil {
			return err
		}

		fmt.Println(helloResponse)
	}

}

// 客户端流 客户端可以发送多次

func LotsOfGreetings(client hello.HelloServiceClient) error {

	// 10秒超时
	ctx, cancelFunc := context.WithDeadline(context.Background(), time.Now().Add(10*time.Second))
	defer cancelFunc()

	greetingsClient, err := client.LotsOfGreetings(ctx)
	if err != nil {
		return err
	}

	defer greetingsClient.CloseSend()

	for i := 0; i < 10; i++ {

		err := greetingsClient.Send(&hello.HelloRequest{
			Name: "客户端流：" + strconv.Itoa(i),
			Age:  18,
		})

		if err != nil {
			return err
		}

	}

	return nil
}

// 双向流

func BidiHello(client hello.HelloServiceClient) error {

	// 10秒超时
	ctx, cancelFunc := context.WithDeadline(context.Background(), time.Now().Add(10*time.Second))
	defer cancelFunc()

	bidiClient, err := client.BidiHello(ctx)
	if err != nil {
		return err
	}

	defer bidiClient.CloseSend()

	for i := 0; i < 10; i++ {

		err := bidiClient.Send(&hello.HelloRequest{
			Name: "双向流" + strconv.Itoa(i),
			Age:  18,
		})

		if err != nil {
			return err
		}

		helloResponse, err := bidiClient.Recv()

		if err == io.EOF {
			return nil
		}

		if err != nil {
			return err
		}

		fmt.Println(helloResponse)
	}

	return nil
}
