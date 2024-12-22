package dao

import (
	"testing"
)

// 删除一条记录
func TestDeleteOneUser(t *testing.T) {
	// User 的 ID 是 75
	user := User{
		ID:       75,
		Username: "wangwu",
	}
	DB.Delete(&user)
	// DELETE FROM `users` WHERE `users`.`id` = 75

	// 带额外条件的删除
	user = User{
		ID:       75,
		Username: "wangwu", // 不生效
	}
	DB.Where("username = ?", "ddda").Delete(&user)
	//  DELETE FROM `users` WHERE username = 'ddda' AND `users`.`id` = 75
}

// 根据主键删除
// GORM 允许通过主键(可以是复合主键)和内联条件来删除对象，它可以使用数字（如以下例子。也可以使用字符串——译者注）
func TestDeleteByPk(t *testing.T) {
	DB.Delete(&User{}, 10)
	//  DELETE FROM `users` WHERE `users`.`id` = 10

	DB.Delete(&User{}, "10")
	//  DELETE FROM `users` WHERE `users`.`id` = '10'

	var users []User
	DB.Delete(&users, []int{1, 2, 3})
	//  DELETE FROM `users` WHERE `users`.`id` IN (1,2,3)
}
