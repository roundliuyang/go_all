package csp

import (
	"fmt"
	"testing"
	"time"
)

func service() string {
	time.Sleep(time.Millisecond * 50)
	return "Done"
}

func otherTask() {
	fmt.Println("working on something else")
	time.Sleep(time.Millisecond * 100)
	fmt.Println("Task is done.")
}

// 串行，浪费时间
func TestService(t *testing.T) {
	fmt.Println(service())
	otherTask()
}

func AsyncService() chan string {
	retCh := make(chan string, 1)
	//retCh := make(chan string, 1)
	go func() {
		ret := service()
		fmt.Println("returned result.")
		retCh <- ret
		fmt.Println("service exited.")
	}()
	return retCh
}

// 通过管道传递消息
func TestAsynService(t *testing.T) {
	retCh := AsyncService()
	otherTask()
	fmt.Println(<-retCh)
	time.Sleep(time.Second * 1)
}
