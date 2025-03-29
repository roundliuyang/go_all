package consumer

import (
	"encoding/json"
	"fmt"
	"github.com/streadway/amqp"
	"rabbitmq-pratice-lesson/conf"
	"rabbitmq-pratice-lesson/lib"
	"rabbitmq-pratice-lesson/model"
	"time"
)

func ConsumeSeller() {
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
	_, err = channel.QueueDeclare(conf.SellerQueue,
		true, false, false, false, nil)
	if err != nil {
		//记录日志
		fmt.Println(err)
		return
	}

	err = channel.QueueBind(conf.SellerQueue, conf.SellerKey, conf.SellerExchange,
		false, nil)
	if err != nil {
		//记录日志
		fmt.Println(err)
		return
	}
	result, err := channel.Consume(conf.SellerQueue, conf.SellerConsumer,
		true, false, false, false, nil)
	if err != nil {
		//记录日志
		fmt.Println(err)
		return
	}

	for item := range result {
		fmt.Println(string(item.Body))
		var order model.Order
		err := json.Unmarshal(item.Body, &order)
		if err != nil {
			//记录日志
			fmt.Println(err)
			continue
			//TODO 如果多次失败，可以进行报警
		}
		if order.ProductId < 1 {
			fmt.Println(order.ID)
			continue
			//TODO 如果多次失败，可以进行报警
		}
		p := model.Product{}
		product, err := p.FindById(order.ProductId)
		if err != nil {
			//记录日志
			fmt.Println(err)
			continue
			//TODO 如果多次失败，可以进行报警
		}
		s := model.Seller{}
		seller := s.FindById(product.SellerId)
		if conf.ProductUp == product.Status && seller.Status == conf.SellerOPEN {
			order.Confirmed = true
			order.Price = product.Price
			sellerProcessOrder(order)
			item.Ack(true)
		} else {
			order.Confirmed = false
		}
	}
}

func sellerProcessOrder(order model.Order) {
	o := model.Order{}
	if order.Confirmed && order.Price > 0 {
		order.Status = model.SellerConfirmed
		b, err := o.Update(model.SellerConfirmed, order.Address, order.ID, order.AccountId,
			order.ProductId, order.DeliveryId, order.CheckOutId, order.PointId, order.Price)
		if err != nil {
			//记录日志
			fmt.Println(err)
			return
		}
		fmt.Println(b)
		sendOrderMessage2DeliverySystem(order)
	} else {
		b, err := o.Update(model.OrderCreateFailed, order.Address, order.ID, order.AccountId,
			order.ProductId, order.DeliveryId, order.CheckOutId, order.PointId, order.Price)
		fmt.Println(b)
		if err != nil {
			//记录日志
			fmt.Println(err)
			return
		}
	}
}

func sendOrderMessage2DeliverySystem(o model.Order) {
	channel, err := lib.GetChannel()
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
	err = channel.Publish(conf.DeliveryExchange, conf.DeliveryKey, false, false, msg)
	if err != nil {
		//记录日志
		fmt.Println(err)
		return
	}
}
