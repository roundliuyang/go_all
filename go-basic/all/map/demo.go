package main

import "fmt"

func main() {

	// 使用var 方式创建
	var m map[string]int
	if m == nil {
		fmt.Println("m is nil")
	}
	fmt.Printf("%p\n", m)
	// 当尝试给nil的map 进行赋值时，就会panic
	//m["age"] = 18

	// new 方式 new(Type) 方式返回的是*Type， 返回的是一个指针，这里的指针已经不是空指针了，是有具体的内存地址了

}
