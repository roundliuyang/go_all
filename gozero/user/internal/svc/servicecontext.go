package svc

import (
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"user/internal/config"
	"user/internal/db"
)

type ServiceContext struct {
	Config config.Config
	Conn   sqlx.SqlConn
}

func NewServiceContext(c config.Config) *ServiceContext {
	sqlConn := db.NewMysql(c.MysqlConfig)
	return &ServiceContext{
		Config: c,
		Conn:   sqlConn,
	}
}
