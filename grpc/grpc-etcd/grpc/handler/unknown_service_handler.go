package handler

import (
	"fmt"
	"google.golang.org/grpc"
	"net/http"
	"sgrpc/grpc/proto/hello"
)

func UnknownServiceHandler(srv interface{}, stream grpc.ServerStream) error {

	fmt.Println("服务未找到...")
	resp := &hello.HelloResponse{
		Code: http.StatusOK,
		Msg:  "not found",
	}

	err := stream.SendMsg(resp)
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}
