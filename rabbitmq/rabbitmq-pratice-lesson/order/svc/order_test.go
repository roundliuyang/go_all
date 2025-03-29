package svc

import (
	"rabbitmq-pratice-lesson/model"
	"testing"
)

func init() {
	model.InitDB()
}

func TestCreateOrder(t *testing.T) {
	status := model.OrderCreating
	address := "北京"
	accountId := 1
	productId := 1
	deliveryId := 0
	settleId := 0
	pointId := 0
	price := 0

	CreateOrder(status, address,
		int64(accountId), int64(productId), int64(deliveryId),
		int64(settleId), int64(pointId), price)
}
