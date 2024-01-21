package dao

import (
	"gorm.io/gorm"
	"log"
	"time"
)

// 定义User模型，绑定users表，ORM库操作数据库，需要定义一个struct类型和MYSQL表进行绑定或者叫映射，struct字段和MYSQL表字段一一对应
type User struct {
	ID int64 // 主键
	//通过在字段后面的标签说明，定义golang字段和表字段的关系
	//例如 `gorm:"column:username"` 标签说明含义是: Mysql表的列名（字段名)为username
	Username string `gorm:"column:username"`
	Password string `gorm:"column:password"`
	//创建时间，时间戳
	CreateTime  int64 `gorm:"column:createtime"`
	UserProfile UserProfile
}

type UserProfile struct {
	ID     int64
	UserId int64
	Sex    int
	Age    int
}

func (u User) TableName() string {
	return "users"
}
func (u UserProfile) TableName() string {
	return "user_profiles"
}

// 插入
func Save(user *User) {
	result := DB.Create(user)
	if result.Error != nil {
		log.Println("insert fail : ", result.Error)
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

// 关联查询========================================================

func Save2() {
	db := DB.Session(&gorm.Session{})
	var user = User{
		Username:   "ms",
		Password:   "ms",
		CreateTime: time.Now().UnixMilli(),
		UserProfile: UserProfile{
			Sex: 0,
			Age: 20,
		},
	}
	db.Save(&user)
}

func GetById2(id int64) User {
	var user User
	err := DB.Preload("UserProfile").Where("id=?", id).First(&user).Error
	err2 := DB.Joins("UserProfile").Where("id=?", id).First(&user).Error
	if err != nil {
		log.Println("query fail : ", err)
	}
	if err2 != nil {
		log.Println("query fail : ", err2)
	}
	return user
}
