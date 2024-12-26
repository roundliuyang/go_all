package dao

import (
	"fmt"
	"strconv"
	"testing"
)

// 查询条件
// 用户的输入只能作为参数，例如：
func TestSqlInjection(t *testing.T) {
	userInput := "jinzhu;drop table users;"

	user := User{}
	// 安全的，会被转义
	DB.Where("username = ?", userInput).First(&user)

	// SQL 注入
	DB.Where(fmt.Sprintf("name = %v", userInput)).First(&user)
}

// 内联条件
func TestSqlInjection2(t *testing.T) {
	userInput := "jinzhu;drop table users;"

	user := User{}
	// 会被转义
	DB.First(&user, "name = ?", userInput)

	// SQL 注入
	DB.First(&user, fmt.Sprintf("name = %v", userInput))

	// 当通过用户输入的整形主键检索记录时，你应该对变量进行类型检查。
	userInputID := "1=1;drop table users;"
	// safe, return error
	id, err := strconv.Atoi(userInputID)
	if err != nil {
		//return err
	}
	DB.First(&user, id)

	// SQL injection
	DB.First(&user, userInputID)
	// SELECT * FROM users WHERE 1=1;drop table users;
}

// SQL 注入方法
// 为了支持某些功能，一些输入不会被转义，调用方法时要小心用户输入的参数。
func TestSqlInjection3(t *testing.T) {
	user := User{}

	DB.Select("name; drop table users;").First(&user)
	DB.Distinct("name; drop table users;").First(&user)

	//DB.Model(&user).Pluck("name; drop table users;", &names)

	DB.Group("name; drop table users;").First(&user)

	DB.Group("name").Having("1 = 1;drop table users;").First(&user)

	DB.Raw("select name from users; drop table users;").First(&user)

	DB.Exec("select name from users; drop table users;")

	DB.Order("name; drop table users;").First(&user)
	// 避免 SQL 注入的一般原则是，不信任用户提交的数据。您可以进行白名单验证来测试用户的输入是否为已知安全的、已批准、已定义的输入，
	// 并且在使用用户的输入时，仅将它们作为参数。
}
