package main

import (
	"flag"
	"fmt"

	"hello01/internal/config"
	"hello01/internal/handler"
	"hello01/internal/svc"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/rest"
)

// 定义配置文件路径
var configFile = flag.String("f", "etc/hello01-api.yaml", "the config file")

func main() {
	flag.Parse()

	// 加载配置
	var c config.Config
	conf.MustLoad(*configFile, &c)

	// rest.RestConf是go-zero提供的配置映射实体，提供了一些默认的配置，方便我们使用。
	server := rest.MustNewServer(c.RestConf)
	defer server.Stop()

	ctx := svc.NewServiceContext(c)
	handler.RegisterHandlers(server, ctx)

	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
	server.Start()
}
