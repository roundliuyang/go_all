package svc

import (
	"fmt"
	"redis-parctice-lesson/06.Points/model"
	"redis-parctice-lesson/global"
	"testing"
)

func init() {
	global.InitGoRedisClient()
	global.InitDB()
	model.InitDB()
}

func TestAddPoints(t *testing.T) {
	nums := GetRandSlice()
	for _, item := range nums {
		fmt.Println(item)
		_, err := AddPoints(int(item), int(item), 1)
		if err != nil {
			panic(err)
		}
	}
}

func TestFindAccountPointsRank(t *testing.T) {
	rank, err := FindAccountPointsRank(0, 9)
	if err != nil {
		panic(t)
	}
	fmt.Println(rank)

	/*
		这个rank里是accountId和积分，用微服务的话，把accountId放到一个切片里，
		然后请求account_srv，拿到对应的用户信息，这样，排行榜的内容就丰富了。
		微服务不在这个课程上讨论，可以参考 从0到。。。
	*/
}
