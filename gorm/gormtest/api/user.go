package api

import (
	"github.com/gin-gonic/gin"
	"gorm/gormtest/dao"
	"time"
)

func SaveUser(c *gin.Context) {
	user := &dao.User{
		Username:   "zhangsan",
		Password:   "123456",
		CreateTime: time.Now().UnixMilli(),
	}
	dao.Save(user)
	c.JSON(200, user)
}

func GetUser(c *gin.Context) {

	user := dao.GetById(1)
	c.JSON(200, user)
}

func GetAll(c *gin.Context) {

	all := dao.GetAll()
	c.JSON(200, all)
}

func UpdateUser(c *gin.Context) {
	dao.UpdateById(1)
	user := dao.GetById(1)
	c.JSON(200, user)
}

func DeleteUser(c *gin.Context) {
	dao.DeleteById(1)
	user := dao.GetById(1)
	c.JSON(200, user)
}
