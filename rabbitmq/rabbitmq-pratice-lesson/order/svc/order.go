package svc

import (
	"encoding/json"
	"fmt"
	"github.com/streadway/amqp"
	"rabbitmq-pratice-lesson/conf"
	"rabbitmq-pratice-lesson/lib"
	"rabbitmq-pratice-lesson/model"
	"time"
)

func CreateOrder(status int, address string,
	accountId, productId, deliverId, checkoutId, pointId int64, price int) {
	o := model.Order{}
	order, err := o.Insert(status, address, accountId, productId, deliverId, checkoutId, pointId, price)
	if err != nil {
		//记录日志
		fmt.Println(err)
		return
	}
	sendMessage2SellerSystem(order)
}

func sendMessage2SellerSystem(o model.Order) {
	channel, err := lib.GetChannel()
	if err != nil {
		//记录日志
		fmt.Println(err)
		return
	}
	err = channel.ExchangeDeclare(conf.SellerExchange, conf.DIRECT,
		true, false, false, false, nil)
	if err != nil {
		//记录日志
		fmt.Println(err)
		return
	}
	orderJson, err := json.Marshal(o)
	if err != nil {
		//记录日志
		fmt.Println(err)
		return
	}
	msg := amqp.Publishing{
		ContentType:  "text/plain",
		DeliveryMode: amqp.Persistent,
		Timestamp:    time.Now(),
		Body:         orderJson,
	}
	err = channel.Publish(conf.SellerExchange, conf.SellerKey, false, false, msg)
	if err != nil {
		//记录日志
		fmt.Println(err)
	}
}
