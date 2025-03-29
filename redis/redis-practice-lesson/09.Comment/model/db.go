package model

import (
	"redis-parctice-lesson/global"
)

func InitDB() {
	err := global.DB.AutoMigrate(&Comment{})
	if err != nil {
		panic("创建Comment表失败:" + err.Error())
	}
}
