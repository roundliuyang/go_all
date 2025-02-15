package main

import (
	"github.com/yly/zap-usage/zap/logwrap/pkg/log"
	"github.com/yly/zap-usage/zap/logwrap/pkg/pkg1"
	"os"
)

func main() {
	file, err := os.OpenFile("./demo1.log", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		panic(err)
	}
	logger := log.New(file, log.InfoLevel)
	log.ResetDefault(logger)
	defer log.Sync()
	log.Info("demo1:", log.String("app", "start ok"),
		log.Int("major version", 2))
	pkg1.Foo()
}
