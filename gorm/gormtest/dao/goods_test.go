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

func TestFindGoods(t *testing.T) {
	FindGoods()
}

func TestFindPageGoods(t *testing.T) {
	FindPageGoods()
}

func TestExecGoods(t *testing.T) {
	ExecGoods()
}

func TestTransactionGoods(t *testing.T) {
	Transaction()
}

func TestTransaction3Goods(t *testing.T) {
	Transaction3()
}
