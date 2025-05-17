package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/nacos-group/nacos-sdk-go/v2/clients"
	"github.com/nacos-group/nacos-sdk-go/v2/common/constant"
	"github.com/nacos-group/nacos-sdk-go/v2/vo"
	"go-frame/cfg"
	"go-frame/knacosregistry"
	"go-frame/proto"
	"go-frame/service"
	grpcstd "google.golang.org/grpc"
	"log"
	"net"
	"os"
	"strconv"
	"time"
)

var (
	logfile string
	quiet   bool

	sconfig *constant.ServerConfig
	cconfig *constant.ClientConfig
)

var (
	Name      = "predis"
	Version   = "unset"
	GitCommit = "unset"
	BuildDate = "unset"
)

func init() {
	initFlags()
	loadnacosenv()
}

func main() {
	initConfig()

	// 创建服务发现客户端（服务发现）
	nc, e := clients.NewNamingClient(vo.NacosClientParam{
		ClientConfig: cconfig, ServerConfigs: []constant.ServerConfig{*sconfig},
	})
	if e != nil {
		log.Fatal(e)
	}

	var userService = service.UserInfoService{}

	// 实例化gRPC
	grpcSrv := grpc.NewServer(
		grpc.Address(":9000"),
		grpc.Middleware(
			recovery.Recovery(),
		),
	)
	//// 在gRPC上注册微服务
	//proto.RegisterUserInfoServiceServer(grpcSrv, &userService)

	app := kratos.New(
		kratos.Name(Name),
		kratos.Version(Version),
		// 服务注册与发现
		kratos.Registrar(knacosregistry.NewRegistry(nc, knacosregistry.WithKind("grpc"))),
		kratos.Server(
			//api.NewHttpServer(),
			grpcSrv,
		),
		kratos.BeforeStop(func(ctx context.Context) error {
			return nil
		}),
	)
	// 在gRPC上注册微服务
	proto.RegisterUserInfoServiceServer(grpcSrv, &userService)

	// 客户端调用 gRPC 测试
	go func() {
		time.Sleep(2 * time.Second)
		callUserInfo()
	}()

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
	log.Printf("Commit:     %s", GitCommit)
	log.Printf("Build Date: %s", BuildDate)
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

func callUserInfo() {
	conn, err := grpcstd.Dial("127.0.0.1:9000", grpcstd.WithInsecure())
	if err != nil {
		log.Printf("Failed to connect to gRPC server: %v", err)
		return
	}
	defer conn.Close()

	client := proto.NewUserInfoServiceClient(conn)

	req := new(proto.UserRequest)
	req.Name = "zs"
	resp, err := client.GetUserInfo(context.Background(), req)
	if err != nil {
		log.Printf("Call GetUserInfo failed: %v", err)
		return
	}

	log.Printf("Call GetUserInfo success: %+v", resp)
}
