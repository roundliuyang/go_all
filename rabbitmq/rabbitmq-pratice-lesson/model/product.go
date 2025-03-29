package model

import (
	"errors"
	"fmt"
	"time"
)

type Product struct {
	ID        int64  `gorm:"id;primary_key"`
	Name      string `gorm:"name"`
	Price     int    `gorm:"price"`
	SellerId  int64  `gorm:"seller_id"`
	Status    int    `gorm:"status"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (p Product) FindById(id int64) (Product, error) {
	var product Product
	//参数校验
	if id < 1 {
		return product, errors.New("参数错误")
	}

	result := DB.First(&product, id)
	if result.Error != nil {
		//TODO生产环境，记录日志
		fmt.Println(result.Error)
	}
	return product, nil
}
