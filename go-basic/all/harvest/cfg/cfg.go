package cfg

import (
	"flag"
	"github.com/spf13/viper"
	"time"
)

type harvestconfig struct {
	Port  string
	Agent struct {
		Cfg string
	}
	Media struct {
		Enabled    bool
		Url        string
		Interval   time.Duration
		MediaItems string
	}
	Taikang struct {
		Punch struct {
			Enabled  bool
			Uid      string
			Url      string
			Interval time.Duration
		}
	}
}

var Cfg = new(harvestconfig)

func init() {
	c := flag.String("c", "app.yml", "配置文件")
	flag.Parse()

	v := viper.New()
	v.SetConfigFile(*c)
	if e := v.ReadInConfig(); e != nil {
		panic(e)
	}
	if e := v.UnmarshalKey("harvest", Cfg); e != nil {
		panic(e)
	}
}
