package dao

import "log"

// 定义User模型，绑定users表，ORM库操作数据库，需要定义一个struct类型和MYSQL表进行绑定或者叫映射，struct字段和MYSQL表字段一一对应
type User struct {
	ID int64 // 主键
	//通过在字段后面的标签说明，定义golang字段和表字段的关系
	//例如 `gorm:"column:username"` 标签说明含义是: Mysql表的列名（字段名)为username
	Username string `gorm:"column:username"`
	Password string `gorm:"column:password"`
	//创建时间，时间戳
	CreateTime int64 `gorm:"column:createtime"`
}

func (u User) TableName() string {
	return "users"
}

// 插入
func Save(user *User) {
	err := DB.Create(user)
	if err != nil {
		log.Println("insert fail : ", err)
	}
}

// 查询
func GetById(id int64) User {
	var user User
	err := DB.Where("id=?", id).First(&user).Error
	if err != nil {
		log.Println("query fail : ", err)
	}
	return user
}

// 查询全部
func GetAll() []User {
	var users []User
	err := DB.Find(&users)
	if err != nil {
		log.Println("get users  fail : ", err)
	}
	return users
}

// 更新
func UpdateById(id int64) {
	err := DB.Model(&User{}).Where("id=?", id).Update("username", "lisisi")
	if err != nil {
		log.Println("update users  fail : ", err)
	}
}

// 删除
func DeleteById(id int64) {
	err := DB.Where("id=?", id).Delete(&User{})
	if err != nil {
		log.Println("delete users  fail : ", err)
	}
}
