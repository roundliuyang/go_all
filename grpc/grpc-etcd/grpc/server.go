package main

import (
	"google.golang.org/grpc"
	"google.golang.org/grpc/health"
	healthpb "google.golang.org/grpc/health/grpc_health_v1"
	"sgrpc/etcd/register"
	"sgrpc/grpc/handler"
	"sgrpc/grpc/proto/hello"
	"sgrpc/grpc/service"

	"log"
	"net"
)

func main() {
	server := grpc.NewServer(grpc.StatsHandler(&handler.StatsHandler{}), grpc.UnknownServiceHandler(handler.UnknownServiceHandler))
	// 健康检查
	healthServer := health.NewServer()
	healthServer.SetServingStatus("grpc.health.v1.Health", healthpb.HealthCheckResponse_SERVING)
	healthpb.RegisterHealthServer(server, healthServer)

	helloService := new(service.HelloService)
	hello.RegisterHelloServiceServer(server, helloService)

	addr := "127.0.0.1:8001"
	listen, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatal(err)
	}

	// --------------------------------
	// 创建一个注册器
	etcdRegister, err := register.NewEtcdRegister()
	if err != nil {
		log.Println(err)
		return
	}
	defer etcdRegister.Close()

	serviceName := "order-service-1"

	// 注册服务
	err = etcdRegister.RegisterServer("/etcd/"+serviceName, addr, 5)
	if err != nil {
		log.Printf("register error %v \n", err)
		return
	}

	// -------------------------------------------------
	server.Serve(listen)
}
