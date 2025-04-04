package service

import (
	"context"
	"fmt"
	"sgrpc/grpc/proto/hello"

	"net/http"
)

type HelloService struct {
	*hello.UnimplementedHelloServiceServer
}

// 一元RPC

func (h HelloService) SayHello(ctx context.Context, request *hello.HelloRequest) (*hello.HelloResponse, error) {

	fmt.Println("say hello")
	msg := "Hello! 8000" + request.GetName()

	resp := &hello.HelloResponse{
		Code: http.StatusOK,
		Msg:  msg,
	}

	return resp, nil
}
