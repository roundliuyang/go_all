package main

import (
	"context"
	"fmt"
	"github.com/openzipkin/zipkin-go"
	zipkingrpc "github.com/openzipkin/zipkin-go/middleware/grpc"
	"github.com/openzipkin/zipkin-go/reporter"
	httpreport "github.com/openzipkin/zipkin-go/reporter/http"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"sgrpc/grpc/proto/hello"
	"time"
)

func main() {

	tracer, r, err := NewZipkinTracer("http://localhost:9411/api/v2/spans", "helloServiceClient", "127.0.0.1:7000")
	defer r.Close()
	if err != nil {
		log.Println(err)
		return
	}

	clientConn, err := grpc.Dial("127.0.0.1:8000", grpc.WithStatsHandler(zipkingrpc.NewClientHandler(tracer)), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		fmt.Println(err)
		return
	}

	helloServiceClient := hello.NewHelloServiceClient(clientConn)

	ctx, cancelFunc := context.WithTimeout(context.Background(), time.Second*5)

	helloResponse, err := helloServiceClient.SayHello(ctx, &hello.HelloRequest{
		Name: "root",
		Age:  10,
	})

	cancelFunc()

	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(helloResponse)
}

// 创建一个zipkin追踪器
func NewZipkinTracer(url, serviceName, hostPort string) (*zipkin.Tracer, reporter.Reporter, error) {

	// 初始化zipkin reporter
	// reporter可以有很多种，如：logReporter、httpReporter，这里我们只使用httpReporter将span报告给http服务，也就是zipkin的http后台
	r := httpreport.NewReporter(url)

	//创建一个endpoint，用来标识当前服务，服务名：服务地址和端口
	endpoint, err := zipkin.NewEndpoint(serviceName, hostPort)
	if err != nil {
		return nil, r, err
	}

	// 初始化追踪器 主要作用有解析span，解析上下文等
	tracer, err := zipkin.NewTracer(r, zipkin.WithLocalEndpoint(endpoint))
	if err != nil {
		return nil, r, err
	}

	return tracer, r, nil
}
