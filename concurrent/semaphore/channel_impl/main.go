package main

import "sync"

/*
实际上，我们还有很多方法实现信号量，比较典型的就是使用 Channel 来实现。
*/
// Semaphore 数据结构，并且还实现了Locker接口
type semaphore struct {
	sync.Locker
	ch chan struct{}
}

// 创建一个新的信号量
func NewSemaphore(capacity int) sync.Locker {
	if capacity <= 0 {
		capacity = 1 // 容量为1就变成了一个互斥锁
	}
	return &semaphore{ch: make(chan struct{}, capacity)}
}

// 请求一个资源
func (s *semaphore) Lock() {
	s.ch <- struct{}{}
}

// 释放资源
func (s *semaphore) Unlock() {
	<-s.ch
}
