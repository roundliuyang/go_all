package model

import (
	"errors"
	"fmt"
	"time"
)

type OrderStatus int

const (
	OrderCreating = iota + 1
	SellerConfirmed
	DeliveryConformed
	CheckoutConformed
	OrderCreated
	OrderCreateFailed
)

type Order struct {
	ID         int64  `gorm:"id;primary_key"`
	Status     int    `gorm:"status"`
	Address    string `gorm:"address"`
	AccountId  int64  `gorm:"account_id"`
	ProductId  int64  `gorm:"product_id"`
	DeliveryId int64  `gorm:"delivery_id"`
	CheckOutId int64  `gorm:"check_out_id"`
	PointId    int64  `gorm:"point_id"`
	Price      int    `gorm:"price"`
	Confirmed  bool   `gorm:"confirmed"`
	CreatedAt  time.Time
	UpdatedAt  time.Time
}

func (o Order) Insert(
	status int, address string,
	accountId, productId, deliveryId, checkOutId, pointId int64,
	price int) (Order, error) {
	order := Order{
		Status:     status,
		Address:    address,
		AccountId:  accountId,
		ProductId:  productId,
		DeliveryId: deliveryId,
		CheckOutId: checkOutId,
		PointId:    pointId,
		Price:      price,
		CreatedAt:  time.Now(),
	}
	result := DB.Save(&order)
	if result.Error != nil {
		//生产环境要记日志
		fmt.Println(result.Error)
		return order, result.Error
	}
	return order, nil
}

func (o Order) Update(status int, address string,
	orderId, accountId, productId, deliveryId, checkOutId, pointId int64,
	price int) (bool, error) {
	if orderId < 1 {
		//记录日志
		return false, errors.New("参数错误")
	}
	var order Order
	DB.First(&order, orderId)
	if order.ID < 1 {
		//记录日志
		return false, errors.New("参数错误")
	}
	result := DB.Model(&order).Updates(
		Order{
			Status:     status,
			Address:    address,
			AccountId:  accountId,
			ProductId:  productId,
			DeliveryId: deliveryId,
			CheckOutId: checkOutId,
			PointId:    pointId,
			Price:      price,
			UpdatedAt:  time.Now(),
		})
	if result.Error != nil {
		return false, result.Error
	}
	if result.RowsAffected > 0 {
		return true, nil
	}
	return false, errors.New("更新失败")
}

func (o Order) FindById(id int64) Order {
	var order Order
	DB.First(&order, id)
	return order
}
