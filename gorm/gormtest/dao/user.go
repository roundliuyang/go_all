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
	CreatedAt   time.Time
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

// Save
func Save(user *User) {
	result := DB.Save(user)
	if result.Error != nil {
		log.Println("save fail : ", result.Error)
	}
}

// 查询------------------------------------------------------------------------------------------------------

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

// 创建---------------------------------------------------------------------------------------------------

func CreateUser(info *User) error {
	err := DB.Create(&info).Error
	return err
}

// 更新----------------------------------------------------------------------------------------------------

// 根据 `struct` 更新属性，只会更新非零值的字段
func UpdateUser(info User) error {
	err := DB.Model(&info).Updates(info).Error
	return err
}

// 根据map更新属性，只会更新非零值的字段
func UpdateUserByMap(info User, param map[string]interface{}) error {
	err := DB.Model(&info).Updates(param).Error
	return err
}

func UpdateUserByMap2(info User, param map[string]interface{}) error {
	err := DB.Model(&info).Where(info).Updates(param).Error
	return err
}

// 删除 ----------------------------------------------------------------------------------------------------

func DeleteById(id int64) {
	err := DB.Where("id=?", id).Delete(&User{})
	if err != nil {
		log.Println("delete users  fail : ", err)
	}
}

// 关联查询-----------------------------------------------------------------------------------------------------

func Save2() {
	db := DB.Session(&gorm.Session{})
	var user = User{
		Username: "ms",
		Password: "ms",
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
