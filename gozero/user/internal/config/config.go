package config

import "github.com/zeromicro/go-zero/rest"

type Config struct {
	rest.RestConf
	MysqlConfig MysqlConfig
	Auth        Auth
}

type MysqlConfig struct {
	DataSource     string
	ConnectTimeout int64
}

type Auth struct {
	Secret string
	Expire int64
}
