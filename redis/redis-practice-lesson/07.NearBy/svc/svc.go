package svc

import (
	"context"
	"errors"
	"github.com/go-redis/redis/v9"
	"redis-parctice-lesson/global"
	"strconv"
)

func UpdateAccountLocation(accountId int, lon, lat float64) error {
	if accountId < 1 {
		return errors.New("参数错误")
	}
	if lon <= 0 {
		return errors.New("获取经度失败")
	}
	if lat <= 0 {
		return errors.New("获取纬度失败")
	}
	key := global.AccountLocation
	geoLocation := &redis.GeoLocation{
		Name:      strconv.Itoa(accountId),
		Longitude: lon,
		Latitude:  lat,
	}
	_, err := global.GoRedisClient.GeoAdd(context.Background(), key, geoLocation).Result()
	return err
}

func FindNearMe(radius int, lon, lat float64) ([]redis.GeoLocation, error) {
	//参数校验

	if radius < 2000 {
		radius = 2000
	}

	key := global.AccountLocation
	query := &redis.GeoRadiusQuery{
		Radius:      float64(radius),
		Unit:        "m",
		WithCoord:   true,
		WithDist:    true,
		WithGeoHash: true,
		Count:       10,
	}
	result, err := global.GoRedisClient.GeoRadius(context.Background(), key, lon, lat, query).Result()
	return result, err
}
