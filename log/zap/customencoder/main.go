package main

import "github.com/yly/zap-usage/zap/customencoder/pkg/log"

func main() {
	log.Info("demo1:", log.String("app", "start ok"))
}
