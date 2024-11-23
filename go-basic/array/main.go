package main

import "fmt"

func main() {
	// 定义一个数组
	original := [3]int{1, 2, 3}

	// 赋值操作
	copy := original

	// 修改副本的值
	copy[0] = 10

	// 打印两个数组
	fmt.Println("Original:", original) // 输出: Original: [1 2 3]
	fmt.Println("Copy:", copy)         // 输出: Copy: [10 2 3]

	fmt.Println("-----------------------------------------------------------")

	var arr1 [5]int
	printArr(&arr1)
	fmt.Println(arr1)
	arr2 := [...]int{2, 4, 6, 8, 10}
	printArr(&arr2)
	fmt.Println(arr2)
}

func printArr(arr *[5]int) {
	arr[0] = 10
	for i, v := range arr {
		fmt.Println(i, v)
	}
}
