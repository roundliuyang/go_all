package model

import (
	"redis-parctice-lesson/global"
)

func InitDB() {
	err := global.DB.AutoMigrate(&Feed{})
	if err != nil {
		panic("创建Feed表失败:" + err.Error())
	}
}
