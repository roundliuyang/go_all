package dao

import (
	"testing"
	"time"
)

func TestSaveGoods(t *testing.T) {
	goods := Goods{
		Title:      "毛巾",
		Price:      25,
		Stock:      100,
		Type:       0,
		CreateTime: time.Now(),
	}
	SaveGoods(goods)
}

func TestUpdateGoods(t *testing.T) {
	UpdateGoods()
}
