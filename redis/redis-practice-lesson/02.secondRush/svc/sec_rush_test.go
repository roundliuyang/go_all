package svc

import (
	"fmt"
	"redis-parctice-lesson/02.secondRush/model"
	"redis-parctice-lesson/global"
	"sync"
	"testing"
)

func init() {
	global.InitDB()
	model.InitDB()
	InitRedis()
	global.InitGoRedisClient()
}

func TestAddVoucher2Redis(t *testing.T) {
	voucherId := 6
	err := AddVoucher2Redis(voucherId)
	if err != nil {
		fmt.Println(err)
	}
}

func TestGetVoucherFromRedis(t *testing.T) {
	voucherId := 6
	result, err := GetVoucherFromRedis(voucherId)
	if err != nil {
		t.Error()
	}
	fmt.Println(*result)
}

func TestDoSecRushV1(t *testing.T) {
	accountId := 1
	voucherId := 6
	err := DoSecRushV1(accountId, voucherId)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println()
}

func TestDoSecRushV1ByGoroutine(t *testing.T) {
	var wg sync.WaitGroup
	for i := 1; i <= 100; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			err := DoSecRushV1(i, 6)
			if err != nil {
				fmt.Println(err)
			} else {
				fmt.Println("DoSecRushV1 suc...")
			}
		}(i)
	}
	wg.Wait()
	fmt.Println("100次抢购结束。。。")
}

func TestDoSecRushV2(t *testing.T) {
	err := DoSecRushV2(6)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("DoSecRushV2 suc...")
	}
	//time.Sleep(20 * time.Second)
}

func TestDoSecRushV4_Simple(t *testing.T) {
	err := DoSecRushV4(6, 6)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("DoSecRushV1 suc...")
	}
}

func TestDoSecRushV4(t *testing.T) {
	var wg sync.WaitGroup
	for i := 1; i <= 1000; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			err := DoSecRushV4(i, 6)
			if err != nil {
				fmt.Println(err)
			} else {
				fmt.Println("DoSecRushV1 suc...")
			}
		}(i)
	}
	wg.Wait()
	fmt.Println("1000次抢购结束。。。")
}
