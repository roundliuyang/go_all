package main

import (
	"fmt"
	"time"
)

/*
使用 Channel 最常见的错误是 panic 和 goroutine 泄漏。
首先，我们来总结下会 panic 的情况，总共有 3 种：
1. close 为 nil 的 chan；
2. send 已经 close 的 chan；
3. close 已经 close 的 chan。
*/
func main() {

}

/*
goroutine 泄漏的问题也很常见，下面的代码也是一个实际项目中的例子：

在这个例子中，process 函数会启动一个 goroutine，去处理需要长时间处理的业务，处理完之后，会发送 true 到 chan 中，
目的是通知其它等待的 goroutine，可以继续处理了。

我们来看一下 x 处，主 goroutine 接收到任务处理完成的通知，或者超时后就返回了。这段代码有问题吗？
如果发生超时，process 函数就返回了，这就会导致 unbuffered 的 chan 从来就没有被读取。我们知道，unbuffered chan 必须等 reader 和 writer
都准备好了才能交流，否则就会阻塞。超时导致未读，结果就是子 goroutine 就阻塞在 y 处永远结束不了，进而导致goroutine 泄漏。
解决这个 Bug 的办法很简单，就是将 unbuffered chan 改成容量为 1 的 chan，这样 y 处就不会被阻塞了。
*/
func process(timeout time.Duration) bool {
	ch := make(chan bool)
	go func() {
		// 模拟处理耗时的业务
		time.Sleep((timeout + time.Second))
		ch <- true // y 处 block

		fmt.Println("exit goroutine")
	}()

	select { // x start
	case result := <-ch:
		return result
	case <-time.After(timeout):
		return false

	} // x end
}

/*
重要：
Go 的开发者极力推荐使用 Channel，不过，这两年，大家意识到，Channel 并不是处理并发问题的“银弹”，有时候使用并发原语更简单，而且不容易出错。
所以，我给你提供一套选择的方法:
1. 共享资源的并发访问使用传统并发原语;
2. 复杂的任务编排和消息传递使用 Channel;
3. 消息通知机制使用 Channel，除非只想 signal 一个 goroutine，才使用 Cond;
4. 简单等待所有任务的完成用 WaitGroup，也有 Channel 的推崇者用 Channel，都可以;
5. 需要和 Select 语句结合，使用 Channel;
6. 需要和超时配合时，使用 Channel 和 Context
*/
