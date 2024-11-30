package main

import (
	"fmt"
	"os"
)

func main() {
	// Fprint-----------------------------------------------------------------------------------
	// 向标准输出写入内容
	fmt.Fprintln(os.Stdout, "向标准输出写入内容!")

	// Sprint系列函数 ----------------------------------------------------------------------------
	s1 := fmt.Sprint("枯藤")
	name := "枯藤"
	age := 18
	s2 := fmt.Sprintf("name:%s,age:%d", name, age)
	s3 := fmt.Sprintln("枯藤")
	fmt.Println(s1, s2, s3)
}
