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

func TestQueryByUserPrimaryKey(t *testing.T) {
	// 如果主键是数字类型，您可以使用 内联条件 来检索对象。 当使用字符串时，需要额外的注意来避免SQL注入
	user := User{}
	DB.First(&user, 10)
	// SELECT * FROM `users` WHERE `users`.`id` = 10 ORDER BY `users`.`id` LIMIT 1

	DB.First(&user, "10")
	fmt.Println(user)

	var users []User
	DB.Find(&users, []int{1, 2, 3})
	//   SELECT * FROM `users` WHERE `users`.`id` IN (1,2,3)

	// 如果主键是字符串(例如像uuid)，查询将被写成如下：
	user2 := User{}
	DB.First(&user2, "id = ?", "1b74413f-f3b8-409f-ac47-e8c062e3472a")
	//  SELECT * FROM `users` WHERE id = '1b74413f-f3b8-409f-ac47-e8c062e3472a' ORDER BY `users`.`id` LIMIT 1

	// 当目标对象有一个主键值时，将使用主键构建查询条件，例如：
	var user3 = User{ID: 10}
	DB.First(&user3)
	//  SELECT * FROM `users` WHERE `users`.`id` = 10 ORDER BY `users`.`id` LIMIT 1

	var result User
	DB.Model(User{ID: 10}).First(&result)
	//  SELECT * FROM `users` ORDER BY `users`.`id` LIMIT 1

	// NOTE: 如果您使用 gorm 的特定字段类型（例如 gorm.DeletedAt），它将运行不同的查询来检索对象。
	var user4 = UserV3{ID: 10}
	DB.First(&user4)
	// SELECT * FROM `users2` WHERE `users2`.`deleted_at` IS NULL AND `users2`.`id` = 10 ORDER BY `users2`.`id` LIMIT 1

}

// 检索全部对象

func TestQueryAllUser(t *testing.T) {
	// Get all records
	user := User{}
	_ = DB.Find(&user)
	//  SELECT * FROM `users`
}

// 条件

// String 条件----------------------------------------------------------------
func TestQueryByStringCondition(t *testing.T) {

	// Get first matched record
	user := User{
		//ID: 1,
	}
	DB.Where("username = ?", "jinzhu").First(&user)
	//SELECT * FROM `users` WHERE username = 'jinzhu' ORDER BY `users`.`id` LIMIT 1   // record not found

	// IN
	var users []User
	DB.Where("username IN ?", []string{"jinzhu", "jinzhu 2"}).Find(&users)
	// SELECT * FROM `users` WHERE username IN ('jinzhu','jinzhu 2')

	// LIKE
	DB.Where("username LIKE ?", "%jin%").Find(&users)
	//  SELECT * FROM `users` WHERE username LIKE '%jin%'

	// AND
	DB.Where("username = ? AND password = ?", "jinzhu", "22").Find(&users)
	// SELECT * FROM `users` WHERE username = 'jinzhu' AND password = '22'

	// BETWEEN

	// 如果对象设置了主键，条件查询将不会覆盖主键的值，而是用 And 连接条件。 例如：

	user = User{
		ID: 10,
	}
	DB.Where("id = ?", 20).First(&user)
	//  SELECT * FROM `users` WHERE id = 20 AND `users`.`id` = 10 ORDER BY `users`.`id` LIMIT 1     // record not found
	// 这个查询将会给出record not found错误 所以，在你想要使用例如 user 这样的变量从数据库中获取新值前，需要将例如 id 这样的主键设置为nil。
}

// Struct & Map 条件----------------------------------------------------------------
func TestQueryByStructOrMapCondition(t *testing.T) {
	// Struct
	user := User{
		//ID: 10,
	}
	DB.Where(&User{Username: "ddda", Password: "123456789"}).First(&user)
	// SELECT * FROM `users` WHERE `users`.`username` = 'ddda' AND `users`.`password` = '123456789' ORDER BY `users`.`id` LIMIT 1

	// Map
	var users []User
	DB.Where(map[string]interface{}{"username": "ddda", "password": "123456789"}).Find(&users)
	// SELECT * FROM `users` WHERE `password` = '123456789' AND `username` =

	// Slice of primary keys
	DB.Where([]int64{20, 21, 22}).Find(&users)
	// SELECT * FROM `users` WHERE `users`.`id` IN (20,21,22)

	// 注意 当使用结构体进行查询时，GORM 将仅使用非零字段进行查询，这意味着如果你的字段的值为 0、''、false 或其他零值，它将不会被用来构建查询条件，例如：
	DB.Where(&User{Username: "ddda", Password: ""}).Find(&users)
	// SELECT * FROM `users` WHERE `users`.`username` = 'ddda'

	// 要在查询条件中包含零值，可以使用映射，它将包含所有键值作为查询条件，例如：
	DB.Where(map[string]interface{}{"username": "ddda", "password": ""}).Find(&users)
	//  SELECT * FROM `users` WHERE `password` = '' AND `username` = 'ddda'
}

// 指定结构体查询字段----------------------------------------------------------------------------

// 使用结构体进行搜索时，可以通过将相关字段名称或 dbname 传递给 Where() 来指定在查询条件中使用结构体中的哪些特定值，例如：
func TestQueryByStructCondition(t *testing.T) {
	var users []User
	DB.Where(&User{Username: "ddda"}, "Password", "Username").Find(&users)
	//  SELECT * FROM `users` WHERE `users`.`username` = 'ddda' AND `users`.`password` = ''
}

// 内联条件---------------------------------------------------------------------------------

// 查询条件可以以与 Where 类似的方式内联到 First 和 Find 等方法中。
func TestQueryByNeiLianCondition(t *testing.T) {
	// Get by primary key if it were a non-integer type
	user := User{}
	DB.First(&user, "id = ?", "75")
	//  SELECT * FROM `users` WHERE id = '75' ORDER BY `users`.`id` LIMIT 1

	// Struct
	var users []User
	DB.Find(&users, User{Username: "ddda"})
	// SELECT * FROM `users` WHERE `users`.`username` = 'ddda'

	// Map
}
