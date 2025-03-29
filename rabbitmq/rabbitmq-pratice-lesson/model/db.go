package model

import (
	"fmt"
	"go.uber.org/zap"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"log"
	"os"
	"rabbitmq-pratice-lesson/global"
	"time"
)

var DB *gorm.DB
var err error

const (
	Port     = 3308
	Name     = "rabbitmq_practice_lesson"
	UserName = "root"
	Password = "123456"
)

//单体服务-viper
//微服务-配置中心

func InitDB() {
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags),
		logger.Config{
			SlowThreshold:             time.Second,
			Colorful:                  true,
			IgnoreRecordNotFoundError: true,
			LogLevel:                  logger.Info,
		},
	)
	conn := fmt.Sprintf("%s:%s@(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		UserName, Password, global.HOST, Port, Name)
	zap.S().Info(conn)
	DB, err = gorm.Open(mysql.Open(conn), &gorm.Config{
		Logger: newLogger,
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
	})
	if err != nil {
		panic("数据库错误:" + err.Error())
	}
	err = DB.AutoMigrate(&Product{}, &Seller{}, &CheckOut{}, &Delivery{}, &Order{}, &Point{})
	if err != nil {
		panic("创建表失败:" + err.Error())
	}
}
