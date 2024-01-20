package main

import (
	"all/all/learn/harvest/cfg"
	"all/all/learn/harvest/taikang"
	"context"
	"github.com/julienschmidt/httprouter"
	"log"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"syscall"
)

var (
	ctx, cancel = context.WithCancel(context.Background())
	wg          sync.WaitGroup
)

func init() {
	go handleSignal()
}
func main() {
	startRoutines()
	router := httprouter.New()
	router.HandlerFunc(http.MethodGet, "/harvest/pingback/record", pingbackRecordHandler)
	// 这个是运动打卡的
	router.Handler(http.MethodGet, "/harvest/api/:version/data", newPbHandler())
	port := cfg.Cfg.Port
	log.Printf("listen and serve on :%s", port)
	log.Fatal(http.ListenAndServe(":"+port, router))
}

func startRoutines() {
	wg.Add(1)
	go func() {
		defer wg.Done()
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		taikang.NewPunchRoutine(cfg.Cfg.Taikang.Punch.Url, cfg.Cfg.Taikang.Punch.Interval).Start(ctx)
	}()

}

// 优雅关闭
func handleSignal() {
	sig := make(chan os.Signal, 10)
	signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)
	log.Printf("signal received :%s", <-sig)

	cancel()
	wg.Wait()
	os.Exit(0)
}
