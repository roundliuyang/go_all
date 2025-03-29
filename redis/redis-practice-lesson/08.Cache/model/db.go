package model

import (
	"redis-parctice-lesson/global"
)

func InitDB() {
	err := global.DB.AutoMigrate(&Product{})
	if err != nil {
		panic("创建Product表失败:" + err.Error())
	}
}
