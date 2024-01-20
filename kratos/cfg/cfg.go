package cfg

import (
	"log"
	"strings"

	"github.com/spf13/viper"
)

var (
	pcfg = new(Cfg)
)

type Cfg struct {
	Concurrent int
	Redis      struct {
		DSN string
	}
	Remote struct {
		Ifcdb    string
		Configdb string
		Excludes []int
		Password string
	}
	Cmsservice struct {
		HotelAllInfo string `mapstructure:"hotel-all-info"`
	}
	WarningUrl string `mapstructure:"warning"`
}

func Init(content string) {
	v := viper.New()
	v.SetConfigType("yaml")
	log.Print("--------------------------------------------------")
	log.Print(content)
	if e := v.ReadConfig(strings.NewReader(content)); e != nil {
		log.Fatalf("read config err: %s", e)
	}

	if e := v.UnmarshalKey("predis", pcfg); e != nil {
		log.Fatalf("unmarshal err: %s", e)
	}
}

func Get() *Cfg {
	return pcfg
}
