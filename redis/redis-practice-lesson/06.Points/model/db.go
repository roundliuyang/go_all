package model

import (
	"redis-parctice-lesson/global"
)

func InitDB() {
	err := global.DB.AutoMigrate(&AccountPoints{})
	if err != nil {
		panic("创建AccountPoints表失败:" + err.Error())
	}
}
