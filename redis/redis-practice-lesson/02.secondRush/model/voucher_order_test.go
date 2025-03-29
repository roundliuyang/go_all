package model

import (
	"fmt"
	"testing"
)

func TestVoucherOrder_Add(t *testing.T) {
	vo := VoucherOrder{}
	voucherId := 6
	accountId := 1
	status := 0
	payment := -1
	orderType := 1
	id, err := vo.Add(voucherId, accountId, status, payment, orderType)
	if err != nil {
		fmt.Println("添加订单失败")
	} else {
		fmt.Printf("添加订单成功%d", id)
	}
}

func TestVoucherOrder_GetVoucherOrderByCondition(t *testing.T) {
	voucherId := 6
	accountId := 1
	vo := VoucherOrder{}

	order, err := vo.GetVoucherOrderByCondition(accountId, voucherId)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(order)
	}
}
