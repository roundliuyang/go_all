package svc

import (
	"context"
	"errors"
	"fmt"
	"github.com/go-redis/redis/v9"
	"redis-parctice-lesson/03.follow_friend/relation"
	"redis-parctice-lesson/04.feeds/model"
	"redis-parctice-lesson/global"
	"strconv"
	"time"
)

type OpMode uint32

const (
	Watch   OpMode = 1
	UnWatch OpMode = 2
)

func CreateFeed(accountId, agreeTotal, commentTotal int, content string) (bool, error) {
	if content == "" || len(content) > 128 {
		return false, errors.New("参数错误")
	}

	feed := model.Feed{}
	id, err := feed.Add(accountId, agreeTotal, commentTotal, 1, content)
	if err != nil {
		return false, err
	}

	followerList, err := relation.FindFollowerList(strconv.Itoa(accountId))
	if err != nil {
		return false, err
	}
	milSecond := time.Now().UnixMilli()
	for _, follower := range followerList {
		key := fmt.Sprintf("%s%s", global.FollowingFeeds, follower)
		global.GoRedisClient.ZAdd(context.Background(), key, redis.Z{
			Score:  float64(milSecond),
			Member: id,
		})
	}
	return true, nil
}

func DeleteFeed(feedId, accountId int) (bool, error) {
	if feedId < 1 {
		return false, errors.New("参数错误")
	}
	f := model.Feed{}
	feed := f.FindById(feedId)
	if feed.ID < 1 {
		return false, errors.New("该Feed已经不存在")
	}
	if feed.AccountId != accountId {
		return false, errors.New("只能删除自己的Feed")
	}
	f.DeletedById(feedId)
	accountIdStr := strconv.Itoa(accountId)
	followers, err := relation.FindFollowerList(accountIdStr)
	if err != nil {
		return false, err
	}
	for _, follower := range followers {
		key := fmt.Sprintf("%s%s", global.FollowingFeeds, follower)
		global.GoRedisClient.ZRem(context.Background(), key, feed.ID)
	}
	return true, nil
}

func ChangeFollowingFeed(followingAccountId, accountId int, kind OpMode) (bool, error) {
	if followingAccountId < 1 {
		return false, errors.New("请选择关注的好友")
	}
	f := model.Feed{}
	feedList := f.FindByAccountId(accountId)
	if len(feedList) < 1 {
		return false, errors.New("参数错误")
	}
	key := fmt.Sprintf("%s%d", global.FollowingFeeds, followingAccountId)
	if kind == UnWatch {
		var ids []interface{}
		for _, feed := range feedList {
			ids = append(ids, feed.ID)
		}
		global.GoRedisClient.ZRem(context.Background(), key, ids...)
	} else {
		now := time.Now()
		for _, feed := range feedList {
			global.GoRedisClient.ZAdd(context.Background(), key, redis.Z{
				Score:  float64(now.UnixMilli()),
				Member: feed.ID,
			})
		}
	}
	return true, nil
}

func SelectForPage(page, accountId int) ([]model.Feed, error) {
	if page < 1 {
		page = 1
	}
	if accountId < 1 {
		return nil, errors.New("参数错误")
	}
	key := fmt.Sprintf("%s%d", global.FollowingFeeds, accountId)
	start := (page - 1) * global.PageSize
	end := page*global.PageSize - 1
	feedIds, err := global.GoRedisClient.ZRevRange(context.Background(), key, int64(start), int64(end)).Result()
	if err != nil {
		return nil, err
	}
	f := model.Feed{}
	feeds := f.FindFeedsByIds(feedIds)
	return feeds, nil
}
