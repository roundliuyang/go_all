package dao

import (
	"fmt"
	"gorm.io/gorm/clause"
	"testing"
)

// 创建记录
func TestCreateUser(t *testing.T) {

	user := User{
		Username: "babalaa2",
		Password: "1111",
	}
	_ = CreateUser(&user)
	fmt.Printf("user id:%d\n", user.ID) // 返回插入数据的主键
	// INSERT INTO `users` (`username`,`password`,`created_at`) VALUES ('babalaa','1111','2024-12-21 23:00:16.313')

	// 创建多项记录
	users := []*User{
		{
			Username: "a",
			Password: "1111",
		},
		{
			Username: "b",
			Password: "1111",
		},
	}
	_ = DB.Create(users)
}

// 默认值

// Upsert 及冲突
func TestUpsertUser(t *testing.T) {

	// Do nothing on conflict
	user := User{
		Username: "upsert",
		Password: "1111",
	}
	DB.Clauses(clause.OnConflict{DoNothing: true}).Create(&user)
	// INSERT INTO `users` (`username`,`password`,`created_at`) VALUES ('upsert','1111','2024-12-21 23:10:58.061') ON DUPLICATE KEY UPDATE `id`=`id`
}
