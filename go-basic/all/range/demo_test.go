package main

import (
	"fmt"
	"testing"
)

func TestSlice(t *testing.T) {
	// 如果我们在遍历数组的同时修改数组的元素，能否得到一个永远都不会停止的循环呢？
	arr := []int{1, 2, 3}
	for _, v := range arr {
		arr = append(arr, v)
	}
	fmt.Println(arr)

	// 神奇的指针
	// 当我们在遍历一个数组时，如果获取 range 返回变量的地址并保存到另一个数组或者哈希时，会遇到令人困惑的现象，下面的代码会输出 “3 3 3”：
	arr = []int{1, 2, 3}
	newArr := []*int{}
	for _, v := range arr {
		newArr = append(newArr, &v)
	}
	for _, v := range newArr {
		fmt.Println(*v)
	}
}

func TestMap(t *testing.T) {
	hash := map[string]int{
		"1": 1,
		"2": 2,
		"3": 3,
	}
	for k, v := range hash {
		println(k, v)
	}
}

func TestPoint(t *testing.T) {
	v := 0
	a := []*int{}

	v = 1
	a = append(a, &v)

	v = 2
	a = append(a, &v)

	v = 3
	a = append(a, &v)

	for _, p := range a {
		fmt.Println(*p)
	}
}
