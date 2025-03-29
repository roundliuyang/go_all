package relation

import (
	"context"
	"errors"
	"redis-parctice-lesson/global"
)

func FindFollowerList(accountId string) ([]string, error) {
	return findCollections(global.Following + accountId)
}

func findCollections(key string) ([]string, error) {
	if key == "" {
		return nil, errors.New("参数错误")
	}
	result, err := global.GoRedisClient.SMembers(context.Background(), key).Result()
	return result, err
}
