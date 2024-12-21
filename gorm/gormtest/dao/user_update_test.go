package dao

import (
	"testing"
)

// Save 会保存所有的字段，即使字段是零值
func TestSaveUser(t *testing.T) {

	user := GetById(7)

	user.Username = "zhangsan"
	user.Password = "123"

	Save(&user)
	// UPDATE `users` SET `username`='zhangsan',`password`='123',`created_at`='2024-02-26 21:40:46' WHERE `id` = 7
}

// 更新单个列
// 当使用 Update 更新单列时，需要有一些条件，否则将会引起ErrMissingWhereClause 错误，查看 阻止全局更新 了解详情。 当使用 Model 方法，并且它有主键值时，主键将会被用于构建条件
func TestSingleColumn(t *testing.T) {
	// 会报错 WHERE conditions required
	user := User{
		//ID:       7,
		Username: "wangwu",
		Password: "",
	}
	DB.Model(&user).Update("username", "hello")

	// 根据model的id更新
	user = User{
		ID:       7,
		Username: "wangwu",
		Password: "",
	}
	DB.Model(&user).Update("username", "zhangsan")
	// UPDATE `users` SET `username`='hello' WHERE `id` = 7

	// 根据条件更新
	DB.Model(&User{}).Where("username = ?", "zhangsan").Update("username", "wangwu")
	// UPDATE `users` SET `username`='wangwu2222' WHERE username = 'wangwu'

	// 根据条件和 model 的值进行更新
	user = User{
		ID:       7,
		Username: "wangwu", // 不生效的
	}
	DB.Model(&user).Where("password = ?", "123").Update("username", "hello")
	// UPDATE `users` SET `username`='hello' WHERE password = '123' AND `id` = 7
}

// 更新多个列
// 注意 使用 struct 更新时, GORM 将只更新非零值字段。 你可能想用 map 来更新属性，或者使用 Select 声明字段来更新
func TestUpdateUser(t *testing.T) {
	user := User{
		ID:       7,
		Username: "wangwu2",
		Password: "",
	}
	_ = UpdateUser(user)
	// UPDATE `users` SET `id`=7,`username`='wangwu' WHERE `id` = 7

	param := map[string]interface{}{
		"username": "wangba",
		"password": "456",
	}
	_ = UpdateUserByMap(user, param)
	// UPDATE `users` SET `password`='456',`username`='wangba' WHERE `id` = 7

	user = User{
		Username: "wangwu2",
		Password: "",
	}
	_ = UpdateUserByMap2(user, param)
	//  UPDATE `users` SET `password`='456',`username`='wangba' WHERE `users`.`username` = 'wangwu2'
}

// 更新或忽略选定字段
func TestSelectFields(t *testing.T) {
	// 选择 Map 的字段
	user := User{
		ID:       7,
		Username: "wangwu2", // 不生效
	}
	DB.Model(&user).Select("username").Updates(map[string]interface{}{"username": "hello", "password": 123456})
	// UPDATE `users` SET `username`='hello' WHERE `id` = 7

	DB.Model(&user).Omit("username").Updates(map[string]interface{}{"username": "hello", "password": 123456})
	// UPDATE `users` SET `password`=123456 WHERE `id` = 7

	// 选择 Struct 的字段（会选中零值的字段）
	DB.Model(&user).Select("Username", "Password").Updates(User{
		ID:       8, // 不生效
		Username: "wangwu2",
		Password: "",
	})
	// UPDATE `users` SET `username`='wangwu2',`password`='' WHERE `id` = 7

	// 选择所有字段（选择包括零值字段的所有字段）,包括主键id也会更新
	DB.Model(&user).Select("*").Updates(User{
		ID:       9,
		Username: "wangwu2",
		Password: "",
	})
	// UPDATE `users` SET `id`=9,`username`='wangwu2',`password`='',`created_at`='0000-00-00 00:00:00' WHERE `id` = 7

	// 选择除 Username 外的所有字段（包括零值字段的所有字段）
	DB.Model(&user).Select("*").Omit("Username").Updates(User{
		ID:       10,
		Username: "wangwu2",
		Password: "",
	})
	// UPDATE `users` SET `id`=10,`password`='',`created_at`='0000-00-00 00:00:00' WHERE `id` = 9
}

// 更新 Hook

// 批量更新

// 阻止全局更新

// 更新的记录数
