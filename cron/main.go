package main

import (
	"fmt"
	"github.com/robfig/cron"
	"log"
	"time"
)

func main() {
	log.Println("Starting...")
	c := cron.New()
	if err := c.AddFunc("* * * * * *", func() { CronFunc() }); err != nil {
		fmt.Errorf("err:%v", err)
		return
	}
	c.Start()
	for {
		select {}
	}
}
func CronFunc() { fmt.Println("CronFunc at:", time.Now()) }
