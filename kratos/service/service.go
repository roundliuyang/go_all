package service

import (
	"context"
	"go-frame/proto"
)

// 定义空接口
type UserInfoService struct {
	proto.UnimplementedUserInfoServiceServer
}

// 实现方法
func (s *UserInfoService) GetUserInfo(ctx context.Context, req *proto.UserRequest) (resp *proto.UserResponse, err error) {
	// 通过用户名查询用户信息
	name := req.Name
	// 数据里查用户信息
	if name == "zs" {
		resp = &proto.UserResponse{
			Id:    1,
			Name:  name,
			Age:   22,
			Hobby: []string{"Sing", "Run"},
		}
	}
	return
}
