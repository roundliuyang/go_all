package account

import (
	"context"
	"errors"
	"time"
	"user/internal/user"

	"user/internal/svc"
	"user/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type RegisterLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RegisterLogic {
	return &RegisterLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RegisterLogic) Register(req *types.RegisterReq) (resp *types.RegisterResp, err error) {
	// todo: add your logic here and delete this line
	userModel := user.NewUserModel(l.svcCtx.Conn)
	u, err := userModel.FindByUsername(l.ctx, req.Username)
	if err != nil {
		l.Logger.Error("Register FindByUsername err: ", err)
		return nil, err
	}
	if u != nil {
		//代表已经注册
		return nil, errors.New("此用户名已经注册")
	}
	_, err = userModel.Insert(l.ctx, &user.User{
		Username:      req.Username,
		Password:      req.Password,
		RegisterTime:  time.Now(),
		LastLoginTime: time.Now(),
	})
	if err != nil {
		return nil, err
	}
	return
	return
}
