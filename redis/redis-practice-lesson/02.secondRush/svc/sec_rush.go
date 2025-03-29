package svc

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/gomodule/redigo/redis"
	"redis-parctice-lesson/02.secondRush/internal"
	"redis-parctice-lesson/02.secondRush/model"
	"redis-parctice-lesson/global"
	"sync"
	"time"
)

var lock sync.Mutex
var rushKeyPrefix = "second_rush_vouchers"

func InitRedis() {
	internal.RedisPool = &redis.Pool{
		MaxIdle:     10,
		MaxActive:   0,
		IdleTimeout: 300 * time.Second,
		Dial: func() (redis.Conn, error) {
			return redis.Dial("tcp", global.HOST+":6389")
		},
	}
}

func GetVoucherFromRedis(voucherId int) (*model.Voucher, error) {
	v := model.Voucher{}
	rushVoucherKey := fmt.Sprintf("SecRush:Voucher:%d", voucherId)
	result, err := global.GoRedisClient.Get(context.Background(), rushVoucherKey).Result()
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal([]byte(result), &v)
	if err != nil {
		return nil, err
	}
	return &v, nil
}

func AddVoucher2Redis(voucherId int) error {
	rv := model.Voucher{}
	result, err := rv.GetById(voucherId)
	if err != nil {
		return err
	}
	r, err := json.Marshal(result)
	if err != nil {
		return err
	}
	rushVoucherKey := fmt.Sprintf("SecRush:Voucher:%d", voucherId)
	b, err := global.GoRedisClient.Set(context.Background(),
		rushVoucherKey, r, time.Hour*24).Result()
	if err != nil {
		return err
	}
	fmt.Println(b)

	rushKey := fmt.Sprintf("SecRush:%d", voucherId)
	stockKey := fmt.Sprintf("SecRush:Stock:%d", voucherId)
	result2, err := global.GoRedisClient.Set(context.Background(), stockKey,
		result.Amount, time.Duration(result.EndTime.UnixMilli())).Result()

	if err != nil {
		return err
	}
	fmt.Println(result2)

	for i := 0; i < result.Amount; i++ {
		result3, err := global.GoRedisClient.RPush(context.Background(), rushKey, 1).Result()
		if err != nil {
			return err
		}
		fmt.Println(result3)
	}

	return nil

}

func DoSecRushV1(accountId, voucherId int) error {
	if voucherId < 1 {
		//记日志
		return errors.New("参数错误")
	}

	if accountId < 1 {
		//记日志
		return errors.New("参数错误")
	}

	rv := model.Voucher{}
	result, err := rv.GetById(voucherId)
	if err != nil {
		return err
	}
	if result.ID < 1 {
		return errors.New("无此活动")
	}

	if result.IsValid < 1 {
		return errors.New("该活动已结束")
	}

	now := time.Now()
	if now.Before(result.StartTime) {
		return errors.New("该抢购还未开始")
	}

	if now.After(result.EndTime) {
		return errors.New("该抢购已经结束")
	}

	if result.Amount < 0 {
		return errors.New("已售罄")
	}

	vo := model.VoucherOrder{}
	resultOrder, err := vo.GetVoucherOrderByCondition(accountId, voucherId)
	if err != nil {
		//可以定义话术返回错误，不暴露设计细节
		return err
	}

	if resultOrder.ID > 0 {
		return errors.New("每人只能抢购1次")
	}

	id, err := rv.DecreaseStock(voucherId)
	if err != nil {
		return err
	}
	if id < 1 {
		return errors.New("扣减库存错误")
	}
	voId, err := vo.Add(voucherId, accountId, 0, -1, 1)
	if err != nil {
		return err
	}

	fmt.Printf("抢购成功%d", voId)
	return nil
}

func DoSecRushV2(voucherId int) error {
	//参数判断
	rv := model.Voucher{}
	result, err := rv.GetById(voucherId)
	if err != nil {
		return err
	}
	//配合redis+锁
	var wg sync.WaitGroup
	//rdb := internal.RedisPool.Get()
	//defer rdb.Close()

	count := result.Amount
	var accountIDList []int
	for i := 1; i <= count; i++ {
		accountIDList = append(accountIDList, i)
	}

	var ch = make(chan int, count)
	rushKey := fmt.Sprintf("%s_%d", rushKeyPrefix, result.ID)
	f := func(accountId int, done func()) {
		defer func() {
			done()
		}()
		lock.Lock()
		defer lock.Unlock()
		lLen, err := global.GoRedisClient.LLen(context.Background(), rushKey).Result()
		if err != nil {
			fmt.Println(err)
		}
		if lLen < 10 {
			do, err := global.GoRedisClient.RPush(context.Background(), rushKey,
				fmt.Sprintf("%d@%v", accountId, time.Now())).Result()
			if err != nil {
				fmt.Println(err)
			}
			fmt.Println(do)
			v := model.Voucher{}
			v.DecreaseStock(voucherId)
			vo := model.VoucherOrder{}
			vo.Add(voucherId, accountId, 0, -1, 1)
			fmt.Printf("%d抢购成功", accountId)
		} else {
			fmt.Printf("%d---抢购活动已结束", accountId)
		}
		ch <- accountId
	}
	wg.Add(count)

	for _, v := range accountIDList {
		go f(v, wg.Done)
	}

	for i := 0; i < count; i++ {
		<-ch
	}
	close(ch)
	wg.Wait()
	return nil
}

func DoSecRushV3(accountIdList []int, voucherId, limited int) error {
	return nil
}

func DoSecRushV4(accountId, voucherId int) error {
	result, err := GetVoucherFromRedis(voucherId)
	if err != nil {
		return err
	}
	//基本的验证要继续做，是否有效，是否时间过期，是否数量为0
	fmt.Println(result)

	rushKey := fmt.Sprintf("SecRush:%d", voucherId)
	r, err := global.GoRedisClient.LPop(context.Background(), rushKey).Result()
	if err != nil && !errors.Is(err, redis.ErrNil) {
		return err
	}
	if r == "1" {
		subStockAndAddOrders(voucherId, accountId)
		fmt.Println(accountId, "抢购成功")
	} else {
		fmt.Println("抢购活动结束")
	}
	return nil
}

func subStockAndAddOrders(voucherId, accountId int) {
	tx := global.DB.Begin()
	var v = model.Voucher{ID: voucherId}
	result := tx.First(&v)
	if result.Error != nil {
		tx.Rollback()
		return
	}
	//可以扣减数据库的amount，也可以扣减redis中的amount
	vo := model.VoucherOrder{
		VoucherID:      voucherId,
		AccountId:      accountId,
		Status:         0,
		Payment:        -1,
		OrderType:      1,
		IsValid:        1,
		CreateDateTime: time.Now(),
	}
	result = tx.Save(&vo)
	if result.Error != nil {
		tx.Rollback()
		return
	}
	tx.Commit()
}
