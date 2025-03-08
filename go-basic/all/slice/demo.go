package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	// 零切片
	slice1 := make([]int, 5)  // 0 0 0 0 0
	slice2 := make([]*int, 5) // nil nil nil nil nil
	fmt.Println(slice1, slice2)

	// nil 切片
	var slice3 []int
	//  只是声明了一个 []int 类型的变量，并未分配底层数组，因此 slice3 的默认值是 nil。
	fmt.Println(slice3 == nil) // true
	fmt.Println(slice3)

	// new([]int) 返回的是 *[]int（指向 nil 切片的指针）。
	// *new([]int) 取消引用得到 nil 切片。
	var slice4 = *new([]int)
	fmt.Println(slice4 == nil) // true

	// 空切片
	var slice5 = []int{}
	var slice6 = make([]int, 0)
	fmt.Println(slice5, slice6)

	type Response struct {
		Items []int `json:"items"`
	}

	// 序列化
	var slice7 []int
	r := Response{
		Items: slice7,
	}
	jsonData, _ := json.Marshal(r)
	fmt.Println(string(jsonData)) // {"items":null}

	var slice8 = make([]int, 0)
	r = Response{
		Items: slice8,
	}
	jsonData, _ = json.Marshal(r)
	fmt.Println(string(jsonData)) // {"items":[]}

	var s []int       // nil
	fmt.Println(s[0]) // panic: index out of range
	//s := make([]int, 0) // 空切片
	//fmt.Println(s[0])   // 仍然 panic，但更明确表示已初始化
}
