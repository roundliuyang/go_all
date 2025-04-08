package main

import "fmt"

type Duck interface {
	Quack()
}

//type Cat struct{}
//
////使用结构体实现接口
//func (c Cat) Quack() {
//	fmt.Println("meow")
//}
//
//func main() {
//	var d1 Duck = Cat{}  //使用结构体初始化变量
//	var d2 Duck = &Cat{} //使用结构体指针初始化变量
//	d1.Quack()
//	d2.Quack()
//}

type Cat struct{}

func (c *Cat) Quack() {
	fmt.Println("meow")
}

func main() {
	//var d1 Duck = Cat{}
	var d2 Duck = &Cat{}
	//d1.Quack()
	d2.Quack()
}
