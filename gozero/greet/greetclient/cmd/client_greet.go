package main

import (
	"context"
	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"greet/greet"
	"log"
)

func main() {
	var clientConf zrpc.RpcClientConf
	conf.MustLoad("greetclient/etc/client.yaml", &clientConf)
	//conn := zrpc.MustNewClient(clientConf)
	conn, err := grpc.NewClient("127.0.0.1:8080", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal(err)
	}
	client := greet.NewGreetClient(conn)
	resp, err := client.Ping(context.Background(), &greet.Request{Ping: "ping"})
	if err != nil {
		log.Fatal(err)
		return
	}
	log.Println(resp)
}
