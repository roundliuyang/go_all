package model

import (
	"redis-parctice-lesson/global"
)

func InitDB() {

	err := global.DB.AutoMigrate(&Voucher{}, &VoucherOrder{})
	if err != nil {
		panic("创建Voucher表失败:" + err.Error())
	}
}
