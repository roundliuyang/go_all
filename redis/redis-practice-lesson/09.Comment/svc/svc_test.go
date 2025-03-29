package svc

import (
	"fmt"
	"redis-parctice-lesson/09.Comment/model"
	"redis-parctice-lesson/global"
	"testing"
)

func init() {
	global.InitDB()
	model.InitDB()
	global.InitGoRedisClient()
}

func TestAddComments(t *testing.T) {
	productId := 100
	content := "从0到Go语言微服务架构师"
	like := 1
	num := 20
	for i := 0; i < num; i++ {
		r, err := AddComments(productId, like, fmt.Sprintf("%s@@@%d", content, i))
		if err != nil {
			panic(err)
		}
		fmt.Println(r)
	}
}

func TestFindLatestComments(t *testing.T) {
	productId := 100
	comments, err := FindLatestComments(productId, 0, 4)
	if err != nil {
		fmt.Println(err)
	}
	for _, item := range comments {
		fmt.Println(item.Content)
	}
}
