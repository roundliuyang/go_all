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

	// 服务上下文 依赖注入，需要用到的依赖都在此进行注入，比如配置，数据库连接，redis连接等
	ctx := svc.NewServiceContext(c)
	// 注册路由
	handler.RegisterHandlers(server, ctx)

	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
	server.Start()
}
