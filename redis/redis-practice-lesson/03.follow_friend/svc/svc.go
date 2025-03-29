package svc

import (
	"context"
	"errors"
	"fmt"
	"redis-parctice-lesson/03.follow_friend/model"
	"redis-parctice-lesson/global"
)

func DoFollow(accountId, followAccountId, isFollowed int) (bool, error) {
	if accountId < 1 || followAccountId < 1 {
		return false, errors.New("参数错误")
	}
	f := model.Follow{}
	follow := f.SelectFollow(accountId, followAccountId)

	if follow.ID < 1 && isFollowed == 1 {
		r, err := f.Add(accountId, followAccountId)
		if err != nil {
			return false, err
		}
		fmt.Println(r)
		result, err := addToRedisSet(accountId, followAccountId)
		if err != nil {
			return false, err
		}
		fmt.Println(result)
		return true, nil
	}
	if follow.ID > 0 && follow.IsValid == 1 && isFollowed == 2 {
		result, err := f.Update(follow.ID, isFollowed)
		if err != nil {
			return false, err
		}
		if result > 0 {
			removeFromRedisSet(accountId, followAccountId)
		}
		return true, nil
	}
	if follow.ID > 0 && follow.IsValid == 2 && isFollowed == 1 {
		result, err := f.Update(follow.ID, isFollowed)
		if err != nil {
			return false, err
		}
		if result > 0 {
			addToRedisSet(accountId, followAccountId)
		}
		return true, nil
	}
	return false, errors.New("未命中操作")
}

func addToRedisSet(accountId, followAccountId int) (int64, error) {
	followingKey := fmt.Sprintf("%s%d", global.Following, accountId)
	r, err := global.GoRedisClient.SAdd(context.Background(), followingKey, followAccountId).Result()
	if err != nil {
		return -1, err
	}
	fmt.Println(r)
	followerKey := fmt.Sprintf("%s%d", global.Followers, followAccountId)
	r, err = global.GoRedisClient.SAdd(context.Background(), followerKey, accountId).Result()
	if err != nil {
		return -1, err
	}
	return 1, nil
}

func removeFromRedisSet(accountId, followAccountId int) (int64, error) {
	followingKey := fmt.Sprintf("%s%d", global.Following, accountId)
	r, err := global.GoRedisClient.SRem(context.Background(), followingKey, followAccountId).Result()
	if err != nil {
		return -1, err
	}
	fmt.Println(r)
	followKey := fmt.Sprintf("%s%d", global.Followers, followAccountId)
	r, err = global.GoRedisClient.SRem(context.Background(), followKey, accountId).Result()
	if err != nil {
		return -1, err
	}
	return 0, nil
}

func FindCommonsFriends(accountId, otherAccountId int) ([]string, error) {
	if accountId < 1 || otherAccountId < 1 {
		return nil, errors.New("参数错误")
	}
	accountFollowingKey := fmt.Sprintf("%s%d", global.Following, accountId)
	otherAccountFollowingKey := fmt.Sprintf("%s%d", global.Following, otherAccountId)
	result, err := global.GoRedisClient.SInter(context.Background(), accountFollowingKey, otherAccountFollowingKey).Result()
	if err != nil {
		return nil, err
	}
	if result == nil || len(result) < 1 {
		return []string{}, nil
	}
	return result, nil
}
