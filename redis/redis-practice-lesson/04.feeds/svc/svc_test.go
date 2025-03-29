package svc

import (
	"fmt"
	"redis-parctice-lesson/03.follow_friend/svc"
	"redis-parctice-lesson/04.feeds/model"
	"redis-parctice-lesson/global"
	"testing"
)

func init() {
	global.InitDB()
	model.InitDB()
	global.InitGoRedisClient()
}

func TestCreateFeed(t *testing.T) {
	accountId := 888
	content := "从0到Go语言微服务架构师" // Go语言微服务架构核心22讲 Go语言极简一本通
	agreeTotal := 99999
	commentTotal := 99999
	r, err := CreateFeed(accountId, agreeTotal, commentTotal, content)
	if err != nil {
		t.Error(err)
	} else {
		fmt.Println(r)
	}
}

func TestDeleteFeed(t *testing.T) {
	//验证只能删除自己的feed这个逻辑分支
	//accountId := 300
	accountId := 888
	b, err := DeleteFeed(6, accountId)
	if !b {
		fmt.Println(err)
	}
}

func TestChangeFollowingFeed(t *testing.T) {
	//关注和取关，要把对应feed添加和删除掉

	followingAccountId := 299
	accountId := 888

	//关注
	b, err := svc.DoFollow(accountId, followingAccountId, 1)
	if err != nil {
		fmt.Println(err)
		return
	}
	if b {
		ChangeFollowingFeed(followingAccountId, accountId, Watch)
	}

	//取关
	//b, err := svc.DoFollow(accountId, followingAccountId, 2)
	//if err != nil {
	//	fmt.Println(err)
	//	return
	//}
	//if b {
	//	ChangeFollowingFeed(followingAccountId, accountId, UnWatch)
	//}
}

func TestSelectForPage(t *testing.T) {
	accountId := 88
	feeds, err := SelectForPage(1, accountId)
	if err != nil {
		t.Error(err)
	} else {
		fmt.Println(feeds)
	}
}
