package model

import (
	"redis-parctice-lesson/global"
)

func InitDB() {

	err := global.DB.AutoMigrate(&Follow{})
	if err != nil {
		panic("创建Follow表失败:" + err.Error())
	}
}
