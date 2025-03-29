package svc

import (
	"fmt"
	"github.com/golang-module/carbon"
	"redis-parctice-lesson/global"
	"sort"
	"testing"
)

func init() {
	global.InitGoRedisClient()
}

func TestDoSign(t *testing.T) {
	accountId := 777

	offset := carbon.Parse("2022-08-01 13:14:15").DayOfMonth()
	//offset := time.Now().Local().Day() - 1
	//offset := 2
	//offset := 5
	//offset := 6
	fmt.Printf("offset1---%d", offset)
	r, err := DoSign(accountId, int64(offset))
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(r)

	offset1 := carbon.Parse("2022-08-05 13:14:15").DayOfMonth()
	fmt.Printf("offset1---%d", offset1)
	r, err = DoSign(accountId, int64(offset1))
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(r)
	offset2 := carbon.Parse("2022-08-06 13:14:15").DayOfMonth()
	fmt.Printf("offset2---%d", offset2)
	r, err = DoSign(accountId, int64(offset2))
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(r)
	offset3 := carbon.Parse("2022-08-07 13:14:15").DayOfMonth()
	fmt.Printf("offset3---%d", offset3)
	r, err = DoSign(accountId, int64(offset3))
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(r)
	offset4 := carbon.Parse("2022-08-08 13:14:15").DayOfMonth()
	fmt.Printf("offset4---%d", offset4)
	r, err = DoSign(accountId, int64(offset4))
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(r)
}

func TestGetSignCount(t *testing.T) {
	accountId := 777
	result, err := GetSignCount(accountId)
	if err != nil {
		t.Error(err)
	}
	fmt.Println(result)
}

func TestGetFirstSignDate(t *testing.T) {
	accountId := 777
	r, err := GetFirstSignDate(accountId)
	if err != nil {
		panic(err)
	}
	fmt.Printf("第一次签到的日期是：%s", r)
}

func TestShowCurrentMonthSignIn(t *testing.T) {
	accountId := 777
	m, err := ShowCurrentMonthSignIn(accountId)
	if err != nil {
		t.Error(err)
	}

	keys := make([]string, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	sort.Sort(sort.StringSlice(keys))
	for _, v := range keys {
		//fmt.Printf("%s:%s\n", v, m[v])
		fmt.Printf("%s:%s\n", v, m[v])
	}
}

func TestGetSignInfo(t *testing.T) {
	accountId := 777
	GetSignInfo(accountId)
}
