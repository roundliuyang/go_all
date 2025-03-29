package svc

import (
	"fmt"
	"redis-parctice-lesson/global"
	"testing"
)

func init() {
	global.InitGoRedisClient()
}

//geoadd：增加某个位置的坐标。
//geopos：获取某个位置的坐标。
//geohash：获取某个位置的geohash值。
//geodist：获取两个位置的距离。
//georadius：根据给定位置坐标获取指定范围内的位置集合。
//georadiusbymember：根据给定位置获取指定范围内的位置集合。

func TestFindNearMe(t *testing.T) {
	lat7 := 39.95231950026053
	lon7 := 116.41422271728516
	result, err := FindNearMe(1000, lon7, lat7)
	if err != nil {
		panic(err)
	}

	fmt.Println(result)
}

func TestUpdateAccountLocation(t *testing.T) {
	lat := 39.91607274850515  //纬度
	lon := 116.40812873840332 //经度

	UpdateAccountLocation(100, lon, lat)

	lat2 := 39.91799827426921  //纬度
	lon2 := 116.39636993408203 //经度

	UpdateAccountLocation(200, lon2, lat2)

	lat3 := 39.936625351474746
	lon3 := 116.38667106628418
	UpdateAccountLocation(300, lon3, lat3)

	lat4 := 39.93391058950708
	lon4 := 116.40379428863525
	UpdateAccountLocation(400, lon4, lat4)
	lat5 := 39.93249557997298
	lon5 := 116.40310764312744
	UpdateAccountLocation(500, lon5, lat5)

	lat6 := 39.93697085890936
	lon6 := 116.40430927276611
	UpdateAccountLocation(600, lon6, lat6)

	lat7 := 39.95231950026053
	lon7 := 116.41422271728516
	UpdateAccountLocation(700, lon7, lat7)
}
