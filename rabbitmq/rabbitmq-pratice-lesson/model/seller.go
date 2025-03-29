package model

import "time"

type Seller struct {
	ID         int64  `gorm:"id:primary_key"`
	Name       string `gorm:name`
	Address    string `gorm:"address"`
	Status     int    `gorm:"status"`
	CheckOutId int64  `gorm:"check_out_id"`
	CreatedAt  time.Time
	UpdatedAt  time.Time
}

func (s Seller) FindById(id int64) Seller {
	var seller Seller
	DB.First(&seller, id)
	return seller
}
