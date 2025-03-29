package model

import (
	"redis-parctice-lesson/global"
	"time"
)

type Comment struct {
	ID         int       `gorm:"primary_key;type:int"`
	ProductId  int       `gorm:"product_id"`
	Content    string    `gorm:"content"`
	Like       int       `gorm:"like"`
	IsValid    int       `gorm:"is_valid"`
	CreateDate time.Time `gorm:"create_date;default:null"`
	UpdateDate time.Time `gorm:"update_date;default:null"`
}

func (c Comment) Add(productId, like int, content string) (Comment, error) {
	comment := Comment{
		ProductId:  productId,
		Content:    content,
		Like:       like,
		IsValid:    1,
		CreateDate: time.Now(),
	}
	result := global.DB.Save(&comment)
	return comment, result.Error
}
