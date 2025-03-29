package svc

import (
	"fmt"
	"redis-parctice-lesson/03.follow_friend/model"
	"redis-parctice-lesson/global"
	"testing"
)

func init() {
	global.InitDB()
	model.InitDB()
	global.InitGoRedisClient()
}

func TestDoFollow(t *testing.T) {
	// follow表，可以先清空

	// 1-添加关注,
	DoFollow(1, 99, 1)
	// 2-取关关注
	DoFollow(1, 99, 2)
	//1-重新关注
	DoFollow(1, 99, 1)

	DoFollow(111, 16, 1)
	DoFollow(111, 127, 1)
	DoFollow(888, 117, 1)
	DoFollow(888, 16, 1)
}

func TestFindCommonsFriends(t *testing.T) {
	accountId := 111
	otherAccountId := 888

	result, err := FindCommonsFriends(accountId, otherAccountId)
	if err != nil {
		t.Error(err)
	}
	fmt.Println(result)
}
