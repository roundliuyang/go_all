package main

import (
	"log"
	"os"
)

func main() {
	// 直接向标准错误(stderr)输出一行日志内容
	log.Println("this is go standard log package")

	// 将日志写入文件
	file, err := os.OpenFile("./demo2.log", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		panic(err)
	}

	// 我们仅需要将实现了io.Writer的os.File传给log包的SetOutput函数即可。这种无需创建logger变量而是直接使用包名+函数的方式写日志的方式减少了传递和管理logger变量的复杂性，这种使用者体验是我们对zap进行封装的目标
	log.SetOutput(file)
	log.Println("this is go standard log package")
}
