package svc

import (
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"hello01/db"
	"hello01/internal/config"
)

type ServiceContext struct {
	Config config.Config
	Conn   sqlx.Session
}

func NewServiceContext(c config.Config) *ServiceContext {
	sqlConn := db.NewMysql(c.MysqlConfig)
	return &ServiceContext{
		Config: c,
		Conn:   sqlConn,
	}
}
