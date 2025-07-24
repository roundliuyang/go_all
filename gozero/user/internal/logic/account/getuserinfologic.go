package account

import (
	"context"
	"database/sql"
	"encoding/json"
	"errors"
	"user/biz"
	"user/internal/user"

	"user/internal/svc"
	"user/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetUserInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserInfoLogic {
	return &GetUserInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetUserInfoLogic) GetUserInfo() (resp *types.UserInfoResp, err error) {
	//如果认证通过 可以从ctx中获取jwt payload
	userId, err := l.ctx.Value("userId").(json.Number).Int64()
	if err != nil {
		return nil, biz.InvalidToken
	}
	u, err := user.NewUserModel(l.svcCtx.Conn).FindOne(l.ctx, userId)
	if err != nil && (errors.Is(err, user.ErrNotFound) ||
		errors.Is(err, sql.ErrNoRows)) {
		return nil, biz.UserNotExist
	}
	resp = &types.UserInfoResp{
		Id:       userId,
		Username: u.Username,
	}
	return
}
