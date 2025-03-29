package main

import (
	"fmt"
	"sync"
	"testing"
)

/*
Channel 的应用场景分为五种类型。这里你先有个印象，这样你可以有目的地去学习 Channel 的基本原理。下节课我会借助具体的例子，来带你掌握这几种类型。
1. 数据交流：当作并发的 buffer 或者 queue，解决生产者 - 消费者问题。多个goroutine 可以并发当作生产者（Producer）和消费者（Consumer）。
2.数据传递：一个 goroutine 将数据交给另一个 goroutine，相当于把数据的拥有权 (引用) 托付出去。
3.信号通知：一个 goroutine 可以将信号 (closing、closed、data ready 等) 传递给另一个或者另一组 goroutine 。
4.任务编排：可以让一组 goroutine 按照一定的顺序并发或者串行的执行，这就是编排的功能。
5. 锁：利用 Channel 也可以实现互斥锁的机制
*/
func dataProducer(ch chan int, wg *sync.WaitGroup) {
	go func() {
		for i := 0; i < 100; i++ {
			ch <- i
		}
		close(ch)

		wg.Done()
	}()

}

func dataReceiver(ch chan int, wg *sync.WaitGroup) {
	go func() {
		for {
			if data, ok := <-ch; ok {
				fmt.Println(data)
			} else {
				break
			}
		}
		wg.Done()
	}()

}

// 如果通道关闭，您仍然可以读取数据。但是您不能将新数据发送到其中
func TestCloseChannel(t *testing.T) {
	var wg sync.WaitGroup
	ch := make(chan int)
	wg.Add(1)
	dataProducer(ch, &wg)
	wg.Add(1)
	dataReceiver(ch, &wg)
	// wg.Add(1)
	// dataReceiver(ch, &wg)
	wg.Wait()
}
