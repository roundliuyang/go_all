package service

import (
	"context"
	klog "github.com/go-kratos/kratos/v2/log"
	"go-frame/proto"
)

// 定义空接口
type UserInfoService struct {
	proto.UnimplementedUserInfoServiceServer
	log *klog.Helper
}

func NewUserInfoService(logger klog.Logger) *UserInfoService {
	return &UserInfoService{
		log: klog.NewHelper(logger),
	}
}

// 实现方法
func (s *UserInfoService) GetUserInfo(ctx context.Context, req *proto.UserRequest) (resp *proto.UserResponse, err error) {
	s.log.Infof("GetUserInfo received, request: %+v", req)
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
