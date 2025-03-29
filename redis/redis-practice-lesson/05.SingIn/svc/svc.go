package svc

import (
	"context"
	"errors"
	"fmt"
	"github.com/go-redis/redis/v9"
	"redis-parctice-lesson/global"
	"time"
)

func DoSign(accountId int, offset int64) (bool, error) {
	signKey := makeSignKey(accountId)
	result, err := global.GoRedisClient.GetBit(context.Background(), signKey, offset).Result()
	if err != nil {
		return false, err
	}
	if result > 0 {
		return true, errors.New("当前日期已经完成签到，无需再签")
	}
	result, err = global.GoRedisClient.SetBit(context.Background(), signKey, offset, 1).Result()
	if err != nil {
		return false, err
	}
	if result == 0 {
		return true, nil
	}
	return false, nil

}

func GetSignCount(accountId int) (int64, error) {
	signKey := makeSignKey(accountId)
	result, err := global.GoRedisClient.BitCount(context.Background(),
		signKey, &redis.BitCount{
			Start: 0,
			End:   10,
		}).Result()
	if err != nil {
		return -1, err
	}
	return result, nil
}

func GetFirstSignDate(accountId int) (string, error) {
	key := makeSignKey(accountId)
	pos, err := global.GoRedisClient.BitPos(context.Background(), key, 1).Result()
	if err != nil {
		return "", err
	}
	pos = pos + 1
	day := time.Now().Local().Day()

	offsetDay := (day - int(pos)) * -1

	return time.Now().AddDate(0, 0, offsetDay).Format("2006-01-02"), nil
}

// ShowCurrentMonthSignIn 获取当月签到情况
func ShowCurrentMonthSignIn(accountId int) (map[string]string, error) {
	signKey := makeSignKey(accountId)
	var day int = time.Now().Local().Day()
	offset := fmt.Sprintf("u%d", day)
	result, err := global.GoRedisClient.BitField(context.Background(),
		signKey, "GET", offset, 0).Result()
	if err != nil {
		return nil, err
	}
	signBoolList := make([]bool, 0)
	signDayList := make([]string, 0)
	v := result[0]
	//fmt.Println(v)
	for i := day; i > 0; i-- {
		pos := (day - i) * -1
		day := time.Now().Local().AddDate(0, 0, pos).Format(global.DateFormat1)
		signDayList = append(signDayList, day)
		var b = v>>1<<1 != v
		signBoolList = append(signBoolList, b)
		v >>= 1
	}
	//fmt.Println(signDayList)
	//fmt.Println(signBoolList)

	m := make(map[string]string)
	for v := 0; v < len(signBoolList); v++ {
		if signBoolList[v] {
			m[signDayList[v]] = "已签到"
		} else {
			m[signDayList[v]] = "未签到"
		}
	}
	return m, nil
}

// GetCurrentMonthSignIn 获取当月签到情况
func GetCurrentMonthSignIn(accountId int) (map[string]string, error) {
	//accountId校验

	// 构建 Key account:sign:7:yyyyMM
	// 0001111100000 //TODO 查一下result[0]这个打印出来的是什么
	signKey := makeSignKey(accountId)

	var day int = time.Now().Local().Day()
	offset := fmt.Sprintf("u%d", day)
	result, _ := global.GoRedisClient.BitField(context.Background(), signKey, "GET", offset, 0).Result()
	signBoolList := make([]bool, 0)
	signDayList := make([]string, 0)
	v := result[0]
	//fmt.Println(v)
	//fmt.Println(v) //这是一个非常大的数字，是2进制转化的10进制数字。
	for i := day; i > 0; i-- {
		pos := (day - i) * -1
		day := time.Now().Local().AddDate(0, 0, pos).Format(global.DateFormat1)
		signDayList = append(signDayList, day)
		var b = v>>1<<1 != v
		signBoolList = append(signBoolList, b)
		v >>= 1
	}

	//fmt.Println(signBoolList)
	//fmt.Println(signDayList)
	m := make(map[string]string)
	for v := 0; v < len(signBoolList); v++ {
		if signBoolList[v] {
			m[signDayList[v]] = "已签到"
		} else {
			m[signDayList[v]] = "未签到"
		}
	}
	//fmt.Println(m)
	return m, nil
}

// 获取当月签到情况
// 根据需要自己实现返回
func GetSignInfo(uid int) (interface{}, error) {
	var keys string = makeSignKey(uid)
	var day int = time.Now().Local().Day()
	var dddd string = fmt.Sprintf("u%d", day)
	st, _ := global.GoRedisClient.BitField(context.Background(), keys, "GET", dddd, 0).Result()
	var res []bool = make([]bool, 0)
	var days []string = make([]string, 0)
	var v int64 = st[0]
	fmt.Println(v)
	for i := day; i > 0; i-- {
		var pos int = (day - i) * -1
		var keys = time.Now().Local().AddDate(0, 0, pos).Format("2006-01-02")
		days = append(days, keys)
		var value = v>>1<<1 != v
		res = append(res, value)
		v >>= 1
	}
	fmt.Println(res)
	fmt.Println(days)
	return nil, nil
}

// account:sign:7:yyyyMM
func makeSignKey(accountId int) string {
	now := time.Now()
	return fmt.Sprintf("account:sign:%d:%s", accountId, now.Format(global.SignFormat))
}
