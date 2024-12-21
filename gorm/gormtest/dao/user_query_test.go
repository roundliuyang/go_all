package dao

import (
	"errors"
	"fmt"
	"gorm.io/gorm"
	"testing"
)

// 检索单个对象

// GORM 提供了 First、Take、Last 方法，以便从数据库中检索单个对象。当查询数据库时它添加了 LIMIT 1 条件，且没有找到记录时，它会返回 ErrRecordNotFound 错误
func TestQueryOneUser(t *testing.T) {
	// 获取第一条记录（主键升序）
	user := User{}
	DB.First(&user)
	// SELECT * FROM `users` ORDER BY `users`.`id` LIMIT 1

	// 获取一条记录，没有指定排序字段
	DB.Take(&user)
	// SELECT * FROM `users` WHERE `users`.`id` = 10 LIMIT 1

	// 获取最后一条记录（主键降序）
	DB.Last(&user)
	// SELECT * FROM `users` WHERE `users`.`id` = 10 ORDER BY `users`.`id` DESC LIMIT 1

	user2 := User{
		ID: 10086,
	}
	result := DB.First(&user2)
	affected := result.RowsAffected
	_ = result.Error
	fmt.Printf("affected:%d\n", affected)
	// 检查 ErrRecordNotFound 错误
	errors.Is(result.Error, gorm.ErrRecordNotFound)

	// 如果你想避免ErrRecordNotFound错误，你可以使用Find，比如db.Limit(1).Find(&user)，Find方法可以接受struct和slice的数据。
	// 对单个对象使用Find而不带limit，db.Find(&user)将会查询整个表并且只返回第一个对象，只是性能不高并且不确定的。
	result = DB.Limit(1).Find(&user2)
	fmt.Println(result.Error)

	fmt.Println("------------------------------------------------------------------------------")
	var user3 User
	//var users3 []User

	// works because destination struct is passed in
	DB.First(&user3)
	//  SELECT * FROM `users` ORDER BY `users`.`id` LIMIT 1

	// works because model is specified using `db.Model()`
	param := map[string]interface{}{}
	DB.Model(&User{}).First(&param)
	// SELECT * FROM `users` ORDER BY `users`.`id` LIMIT 1

	// doesn't work
	param2 := map[string]interface{}{}
	DB.Table("users").First(&param2)

	// works with Take
	param3 := map[string]interface{}{}
	DB.Table("users").Take(&param3)
	// SELECT * FROM `users` LIMIT 1

	// no primary key defined, results will be ordered by first field (i.e., `Code`)
	type User struct {
		Username string
	}
	DB.First(&User{})
	//  SELECT * FROM `users` ORDER BY `users`.`username` LIMIT 1
}

// 根据主键检索
