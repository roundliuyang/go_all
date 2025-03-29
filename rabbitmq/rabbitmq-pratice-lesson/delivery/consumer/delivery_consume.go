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

func ConsumeDelivery() {
	channel, err := lib.GetChannel()
	if err != nil {
		//记录日志
		fmt.Println(err)
		return
	}
	err = channel.ExchangeDeclare(conf.DeliveryExchange, conf.DIRECT,
		true, false, false, false, nil)
	if err != nil {
		//记录日志
		fmt.Println(err)
		return
	}
	_, err = channel.QueueDeclare(conf.DeliveryQueue,
		true, false, false, false, nil)
	if err != nil {
		//记录日志
		fmt.Println(err)
		return
	}
	err = channel.QueueBind(conf.DeliveryQueue, conf.DeliveryKey, conf.DeliveryExchange, false, nil)
	if err != nil {
		//记录日志
		fmt.Println(err)
		return
	}
	result, err := channel.Consume(conf.DeliveryQueue, conf.DeliveryConsumer, true, false, false, false, nil)
	if err != nil {
		//记录日志
		fmt.Println(err)
		return
	}
	for item := range result {
		var order model.Order
		err := json.Unmarshal(item.Body, &order)
		if err != nil {
			//记录日志,设置报警环节
			fmt.Println(err)
			continue
		}
		d := model.Delivery{}
		deliveryList, err := d.FindByStatus(model.DeliveryStatusFree)
		order.DeliveryId = deliveryList[0].ID
		deliveryProcessOrder(order)
		item.Ack(true)
	}
}

func deliveryProcessOrder(order model.Order) {
	o := model.Order{}
	if order.Status == model.SellerConfirmed && order.DeliveryId > 0 {
		order.Status = model.DeliveryConformed
		o.Update(model.DeliveryConformed, order.Address, order.ID,
			order.AccountId, order.ProductId, order.DeliveryId, order.CheckOutId, order.PointId, order.Price)
		sendOrderMessage2CheckOutSystem(order)
	}
}

func sendOrderMessage2CheckOutSystem(o model.Order) {
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
	err = channel.Publish(conf.CheckoutExchange, conf.CheckoutKey, false, false, msg)
	if err != nil {
		//记录日志
		fmt.Println(err)
	}
}
