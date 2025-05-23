package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/go-kratos/kratos/v2"
	klog "github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/middleware/tracing"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/nacos-group/nacos-sdk-go/v2/clients"
	"github.com/nacos-group/nacos-sdk-go/v2/common/constant"
	"github.com/nacos-group/nacos-sdk-go/v2/vo"
	"go-frame/api"
	"go-frame/cfg"
	"go-frame/knacosregistry"
	"go-frame/proto"
	"go-frame/service"
	"go-frame/tracer"
	"go.opentelemetry.io/otel"
	"log"
	"net"
	"os"
	"strconv"
)

var (
	logfile string
	quiet   bool

	sconfig *constant.ServerConfig
	cconfig *constant.ClientConfig
)

var (
	Name    = "predis"
	Version = "unset"

	id, _ = os.Hostname()
)

func init() {
	initFlags()
	loadnacosenv()
}

func main() {
	initConfig()

	// 配置，启动链路追踪
	url := "http://172.26.118.30:14268/api/traces"
	Name = "kratos.service.predis"
	id = "kratos.id.user.1"
	Version = "test-V0.0.1"
	traceconf := tracer.NewConf(Name, id, Version, url)
	tp, _ := traceconf.TracerProvider()
	otel.SetTracerProvider(tp) // 为全局链路追踪

	logger := klog.With(klog.NewStdLogger(os.Stdout),
		"ts", klog.DefaultTimestamp,
		"caller", klog.DefaultCaller,
		"service.id", id,
		"service.name", Name,
		"service.version", Version,
		"trace.id", tracing.TraceID(),
		"span.id", tracing.SpanID(),
	)

	// 创建服务发现客户端（服务发现）
	nc, e := clients.NewNamingClient(vo.NacosClientParam{
		ClientConfig: cconfig, ServerConfigs: []constant.ServerConfig{*sconfig},
	})
	if e != nil {
		log.Fatal(e)
	}

	userService := service.NewUserInfoService(logger)

	// 实例化gRPC
	grpcSrv := grpc.NewServer(
		grpc.Address(":9000"),
		grpc.Middleware(
			middleware.Chain(
				recovery.Recovery(),
				tracing.Server(), //设置trace
			),
		),
	)
	// 在gRPC上注册微服务
	proto.RegisterUserInfoServiceServer(grpcSrv, userService)

	app := kratos.New(
		kratos.Name(Name),
		kratos.Version(Version),
		// 服务注册与发现
		kratos.Registrar(knacosregistry.NewRegistry(nc, knacosregistry.WithKind("grpc"))),
		kratos.Logger(logger),
		kratos.Server(
			api.NewHttpServer(),
			grpcSrv,
		),
		kratos.BeforeStop(func(ctx context.Context) error {
			return nil
		}),
	)

	if e := app.Run(); e != nil {
		log.Fatal(e)
	}
}

func loadnacosenv() {
	//addr := os.Getenv("NACOS_ADDR")
	//username := os.Getenv("NACOS_USERNAME")
	//password := os.Getenv("NACOS_PASSWORD")
	//namespace := os.Getenv("NACOS_NS")

	addr := "172.26.118.30:8848"
	username := "nacos"
	password := "nacos"
	namespace := "idc"

	// 创建clientConfig
	cconfig = constant.NewClientConfig(func(cc *constant.ClientConfig) {
		cc.NamespaceId = namespace
		cc.Username = username
		cc.Password = password
	})
	ipaddr, port, e := net.SplitHostPort(addr)
	if e != nil {
		log.Panicf("parse NACOS_ADDR err: %s", e)
	}
	iport, e := strconv.ParseUint(port, 10, 64)
	if e != nil {
		log.Panicf("parse port err: %s", e)
	}
	// serverConfig
	sconfig = constant.NewServerConfig(ipaddr, iport)
}

func initConfig() {
	// 创建动态配置客户端（动态配置）
	cc, e := clients.NewConfigClient(vo.NacosClientParam{
		ClientConfig:  cconfig,
		ServerConfigs: []constant.ServerConfig{*sconfig},
	})
	if e != nil {
		log.Fatal(e)
	}
	p, e := cc.GetConfig(vo.ConfigParam{
		DataId: "predis.yml",
	})
	if e != nil {
		log.Fatal(e)
	}
	cfg.Init(p)

	if cc.ListenConfig(vo.ConfigParam{
		Group: constant.DEFAULT_GROUP, DataId: "predis.yml",
		OnChange: func(namespace, group, dataId, data string) {

			cfg.Init(data)
			run()
		},
	}); e != nil {
		log.Panicf("watch config err: dataId = predis.example.yml, err = %s", e)
	}
}

func run() {
	log.Printf("Name:       predis")
	log.Printf("Version:    %s", Version)
	log.Printf("Config: 	%+v", cfg.Get())
}

func initFlags() {
	help := flag.Bool("h", false, "打印帮助信息")
	flag.StringVar(&logfile, "log", "/data/logs/predis/predis.log", "日志")
	flag.BoolVar(&quiet, "q", false, "不输出到stdout")

	flag.Usage = func() {
		fmt.Println("Usage of predis: ")
		flag.PrintDefaults()
		os.Exit(0)
	}
	flag.Parse()

	if *help || logfile == "" {
		flag.Usage()
	}
}
