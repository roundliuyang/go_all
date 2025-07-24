package main

import (
	"flag"
	"fmt"

	"greet/greet"
	"greet/internal/config"
	"greet/internal/server"
	"greet/internal/svc"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/service"
	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var configFile = flag.String("f", "etc/greet.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)
	ctx := svc.NewServiceContext(c)

	s := zrpc.MustNewServer(c.RpcServerConf, func(grpcServer *grpc.Server) {
		// 注册 Greet 服务到 gRPC 服务器中。server.NewGreetServer(ctx) 创建一个新的 GreetServer 实例，它通常包含了业务逻辑，负责处理客户端的 RPC 请求
		greet.RegisterGreetServer(grpcServer, server.NewGreetServer(ctx))

		if c.Mode == service.DevMode || c.Mode == service.TestMode {
			// 注册反射服务，在gRPC中，反射是一种机制，允许客户端在不知道服务定义（即.proto文件）的情况下查询服务端上的gRPC服务信息。
			reflection.Register(grpcServer)
		}
	})
	defer s.Stop()

	fmt.Printf("Starting rpc server at %s...\n", c.ListenOn)
	s.Start()
}
