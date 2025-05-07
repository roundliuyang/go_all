package main

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// 定义模型
type Product struct {
	gorm.Model
	Code  string
	Price uint
}

func main() {
	dsn := "root:123456@tcp(127.0.0.1:3306)/test?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// 自动迁移模式
	db.AutoMigrate(&Product{})

	// 插入3条示例数据
	db.Create(&Product{Code: "D42", Price: 100})
	db.Create(&Product{Code: "D43", Price: 200})
	db.Create(&Product{Code: "D44", Price: 300})

	// 根据不同条件查询
	var products []Product

	// 看下面一段代码，其中的三次查询为何不会相会影响，查询一的Where不会带到查询二和查询三。
	// 查询一
	db.Where("price > ?", 150).Find(&products) // 查询价格大于150的产品
	fmt.Println("Products with price > 150:", products)

	// 查询二
	db.Where("code LIKE ?", "%42%").Find(&products) // 查询代码包含42的产品
	fmt.Println("Products with code containing 42:", products)

	// 查询三
	db.Where("price BETWEEN ? AND ?", 100, 200).Find(&products) // 查询价格在100到200之间的产品
	fmt.Println("Products with price between 100 and 200:", products)
}
