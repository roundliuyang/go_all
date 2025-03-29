package model

import (
	"fmt"
	"redis-parctice-lesson/global"
	"time"
)

var (
	UserStatusOK  = 1
	UserStatusDel = 10
)

type User struct {
	ID     int64     `json:"id" gorm:"type BIGINT(32);primaryKey"`
	Name   string    `json:"name" gorm:"column:name;type:VARCHAR(50);default ''"`
	Passwd string    `json:"passwd" gorm:"column:passwd;type:VARCHAR(512)"`
	Email  string    `json:"email" gorm:"column:email;type:VARCHAR(100)"`
	Mobile string    `json:"mobile" gorm:"column:mobile;type:VARCHAR(20)"`
	Salt   string    `json:"salt" gorm:"column:salt;type:VARCHAR(4)"`
	Status int       `json:"status" gorm:"column:status;type:TINYINT(4)"`
	Ctime  time.Time `json:"ctime" gorm:"column:ctime;type:TIMESTAMP;default 'CURRENT_TIMESTAMP'"`
	Mtime  time.Time `json:"mtime" gorm:"column:mtime;type:TIMESTAMP;"`
}

func (u *User) GetRow(id int64) User {
	var user User
	global.DB.Model(User{}).Where("id=?", id).Find(&user)
	return user
}

func (u *User) GetAll() []User {
	var users []User
	global.DB.Model(User{}).Find(&users)
	return users
}

func (u *User) Add(email, mobile, name, passwd, salt string) User {
	now := time.Now()

	user := User{
		Name:   name,
		Passwd: passwd,
		Email:  email,
		Mobile: mobile,
		Salt:   salt,
		Status: UserStatusOK,
		Ctime:  now,
		Mtime:  now,
	}

	result := global.DB.Save(&user)
	if result.Error != nil {
		//生产环境，要记录日志
		fmt.Println(result.Error)
	}
	return user
}

func IsExistsMobile(mobile string) User {
	var user User
	global.DB.Model(User{}).Where("mobile=?", mobile).Find(&user)
	return user
}

func (u User) GetRowById(id int64) User {
	var user User
	global.DB.First(&user, id)
	return user
}
