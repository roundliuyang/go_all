package dao

import (
	"fmt"
	"gorm.io/gorm/clause"
	"testing"
	"time"
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

// 默认值 详见：https://gorm.io/zh_CN/docs/create.html
func TestDefaultValue(t *testing.T) {
	user := User{
		Username: "",
	}
	_ = CreateUser(&user)
	// INSERT INTO `users` (`username`,`password`,`created_at`) VALUES ('ddda','123456789','2024-12-22 00:43:51.753')

	// 注意，当结构体的字段默认值是零值的时候（如上User）比如 0, '', false，这些字段值将不会被保存到数据库中，你可以使用指针类型或者Scanner/Valuer来避免这种情况。
	username := "example_user"
	password := ""
	userV2 := UserV2{
		Username:  &username, // 指向非零值的指针
		Password:  &password, // password 字符串的零值会被保存到数据库
		CreatedAt: time.Now(),
	}
	_ = CreateUserV2(&userV2)
	//  INSERT INTO `users` (`username`,`password`,`created_at`) VALUES ('example_user','','2024-12-22 00:55:55.653')
}

// Upsert 及冲突
func TestUpsertUser(t *testing.T) {

	// Do nothing on conflict
	user := User{
		Username: "upsert",
		Password: "1111",
	}
	DB.Clauses(clause.OnConflict{DoNothing: true}).Create(&user)
	// INSERT INTO `users` (`username`,`password`,`created_at`) VALUES ('upsert','1111','2024-12-21 23:10:58.061') ON DUPLICATE KEY UPDATE `id`=`id`

	// Update all columns to new value on conflict except primary keys and those columns having default values from sql func
	user2 := User{
		Username: "upsert",
		Password: "abc",
	}
	DB.Clauses(clause.OnConflict{
		UpdateAll: true,
	}).Create(&user2)
	// INSERT INTO `users` (`username`,`password`,`created_at`) VALUES ('upsert','abc','2024-12-21 23:44:00.399') ON DUPLICATE KEY UPDATE `username`=VALUES(`username`),`password`=VALUES(`password`)

	// Update columns to new value on `id` conflict
	user3 := User{
		Username: "upsert",
		Password: "888888888888888888888",
	}
	DB.Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "id"}}, // TMD，指定了id,其他唯一约束冲突照样会更新
		DoUpdates: clause.AssignmentColumns([]string{"username", "password"}),
	}).Create(&user3)
	// INSERT INTO `users` (`username`,`password`,`created_at`) VALUES ('upsert','888888888888888888888','2024-12-22 00:05:24.451') ON DUPLICATE KEY UPDATE `username`=VALUES(`username`),`password`=VALUES(`password`)
}

// 关联创建
