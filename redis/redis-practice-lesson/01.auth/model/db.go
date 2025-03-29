package model

import (
	"redis-parctice-lesson/global"
)

func InitDB() {
	err := global.DB.AutoMigrate(&User{})
	if err != nil {
		panic("创建User表失败:" + err.Error())
	}
}
