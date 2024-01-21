package dao

import (
	"testing"
	"time"
)

func TestSaveUser(t *testing.T) {
	user := &User{
		Username:   "zhangsan",
		Password:   "123456",
		CreateTime: time.Now().UnixMilli(),
	}
	Save(user)
}

func TestSave2User(t *testing.T) {
	Save2()
}

func TestGetById2User(t *testing.T) {
	GetById2(10)
}
