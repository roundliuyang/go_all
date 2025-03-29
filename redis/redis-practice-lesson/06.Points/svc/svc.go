package svc

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v9"
	"math/rand"
	"redis-parctice-lesson/06.Points/model"
	"redis-parctice-lesson/global"
	"strconv"
	"time"
)

func AddPoints(accountId, points, kind int) (float64, error) {
	p := model.AccountPoints{}
	p.Add(accountId, points, kind)
	accountIdStr := strconv.Itoa(accountId)
	key := global.AccountPoints
	result, err := global.GoRedisClient.ZIncrBy(context.Background(), key, float64(points), accountIdStr).Result()
	if err != nil {
		return -1, err
	}
	return result, nil
}

func FindAccountPointsRank(start, end int64) ([]redis.Z, error) {
	key := global.AccountPoints
	result, err := global.GoRedisClient.ZRevRangeWithScores(context.Background(), key, start, end).Result()
	if err != nil {
		return nil, err
	}
	if result == nil || len(result) < 1 {
		return nil, nil
	}
	for _, item := range result {
		fmt.Sprintf("%s--%f", item.Member, item.Score)
	}
	return result, nil
}

func GetRandSlice() []int64 {
	var s []int64
	rand.Seed(time.Now().UnixNano()) // 纳秒时间戳
	for i := 1; i <= 1000; i++ {
		data := rand.Int63n(1000)
		s = append(s, data)
	}
	return s
}
