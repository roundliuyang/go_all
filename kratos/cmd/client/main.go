package main

import (
	"context"
	"github.com/go-kratos/kratos/contrib/registry/nacos/v2"
	"github.com/go-kratos/kratos/v2/middleware/tracing"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/nacos-group/nacos-sdk-go/clients"
	"github.com/nacos-group/nacos-sdk-go/common/constant"
	"github.com/nacos-group/nacos-sdk-go/vo"
	"go-frame/proto"
	"log"
)

func main() {
	sc := []constant.ServerConfig{
		*constant.NewServerConfig("172.26.118.30", 8848),
	}

	cc := &constant.ClientConfig{
		NamespaceId:         "idc",
		TimeoutMs:           5000,
		NotLoadCacheAtStart: true,
		//LogDir:              "/tmp/nacos/log",
		//CacheDir:            "/tmp/nacos/cache",
		LogLevel: "debug",
		Username: "nacos",
		Password: "nacos",
	}

	cli, err := clients.NewNamingClient(
		vo.NacosClientParam{
			ClientConfig:  cc,
			ServerConfigs: sc,
		},
	)
	if err != nil {
		log.Panic(err)
	}

	// 构造 gRPC 客户端连接，通过 Nacos 自动发现服务地址
	conn, err := grpc.DialInsecure(
		context.Background(),
		grpc.WithEndpoint("discovery:///kratos.service.predis"),
		grpc.WithDiscovery(nacos.New(cli)),
		grpc.WithMiddleware( // ✅ tracing 客户端中间件
			tracing.Client(),
		),
	)
	client := proto.NewUserInfoServiceClient(conn)
	if err != nil {
		log.Fatal("grpc dial error: ", err)
	}
	defer conn.Close()

	req := new(proto.UserRequest)
	req.Name = "zs"
	resp, err := client.GetUserInfo(context.Background(), req)
	if err != nil {
		log.Printf("Call GetUserInfo failed: %v", err)
		return
	}

	log.Printf("Call GetUserInfo success: %+v", resp)
}

//func callUserInfo() {
//	conn, err := grpcstd.Dial("127.0.0.1:9000", grpcstd.WithInsecure())
//	if err != nil {
//		log.Printf("Failed to connect to gRPC server: %v", err)
//		return
//	}
//	defer conn.Close()
//
//	client := proto.NewUserInfoServiceClient(conn)
//
//	req := new(proto.UserRequest)
//	req.Name = "zs"
//	resp, err := client.GetUserInfo(context.Background(), req)
//	if err != nil {
//		log.Printf("Call GetUserInfo failed: %v", err)
//		return
//	}
//
//	log.Printf("Call GetUserInfo success: %+v", resp)
//}
