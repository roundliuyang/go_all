package model

import (
	"fmt"
	"redis-parctice-lesson/global"
	"testing"
	"time"
)

func init() {
	global.InitDB()
	InitDB()
	global.InitGoRedisClient()
}

func TestVoucher_Add(t *testing.T) {
	now := time.Now()
	amount := 100
	start := now
	endTime := now.AddDate(0, 0, 10)
	rv := &Voucher{}
	id, err := rv.Add(amount, start, endTime)
	if err != nil {
		fmt.Println("添加失败")
	} else {
		fmt.Printf("添加成功ID:%d", id)
	}
}

func TestVoucher_GetById(t *testing.T) {
	id := 6
	rv := &Voucher{}
	voucher, err := rv.GetById(id)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(voucher)
}

func TestVoucher_DecreaseStock(t *testing.T) {
	voucherId := 6
	rv := &Voucher{}
	result, err := rv.DecreaseStock(voucherId)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(result)
}
