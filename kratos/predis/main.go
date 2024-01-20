package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/go-kratos/kratos/v2"
	"github.com/nacos-group/nacos-sdk-go/v2/clients"
	"github.com/nacos-group/nacos-sdk-go/v2/common/constant"
	"github.com/nacos-group/nacos-sdk-go/v2/vo"
	"go-frame/api"
	"go-frame/cfg"
	"go-frame/knacosregistry"
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
	nc, e := clients.NewNamingClient(vo.NacosClientParam{
		ClientConfig: cconfig, ServerConfigs: []constant.ServerConfig{*sconfig},
	})
	if e != nil {
		log.Fatal(e)
	}

	app := kratos.New(
		kratos.Name(Name),
		kratos.Version(Version),
		kratos.Registrar(knacosregistry.NewRegistry(nc, knacosregistry.WithKind("http"))),
		kratos.Server(api.NewHttpServer()),
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

	addr := "172.19.171.168:8848"
	username := "nacos"
	password := "nacos"
	namespace := "idc"

	cconfig = constant.NewClientConfig(func(cc *constant.ClientConfig) {
		cc.NamespaceId = namespace
		cc.Username = username
		cc.Password = password
	})
	ipaddr, port, e := net.SplitHostPort(addr)
	if e != nil {
		// todo 区别
		log.Panicf("parse NACOS_ADDR err: %s", e)
	}
	iport, e := strconv.ParseUint(port, 10, 64)
	if e != nil {
		log.Panicf("parse port err: %s", e)
	}
	sconfig = constant.NewServerConfig(ipaddr, iport)
}

func initConfig() {
	cc, e := clients.NewConfigClient(vo.NacosClientParam{
		ClientConfig:  cconfig,
		ServerConfigs: []constant.ServerConfig{*sconfig},
	})
	if e != nil {
		log.Fatal("init nacos failed!")
		log.Fatal(cconfig)
		log.Fatal(e)
	}
	p, e := cc.GetConfig(vo.ConfigParam{
		DataId: "predis.yml",
	})
	if e != nil {
		log.Fatal("get nacos predis.yml faild!")
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
		log.Panicf("watch config err: dataId = predis.yml, err = %s", e)
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
