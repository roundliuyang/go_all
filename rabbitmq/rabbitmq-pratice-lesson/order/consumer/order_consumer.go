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

func ConsumeOrder() {
	channel, err := lib.GetChannel()
	if err != nil {
		//记录日志
		fmt.Println(err)
		return
	}

	err = channel.ExchangeDeclare(conf.OrderExchange, conf.TOPIC, true,
		false, false, false, nil)
	if err != nil {
		//记录日志
		fmt.Println(err)
		return
	}

	_, err = channel.QueueDeclare(conf.OrderQueue, true,
		false, false, false, nil)
	if err != nil {
		//记录日志
		fmt.Println(err)
		return
	}

	err = channel.QueueBind(conf.OrderQueue, conf.OrderConsumeKey, conf.OrderExchange, false, nil)
	if err != nil {
		//记录日志
		fmt.Println(err)
		return
	}
	result, err := channel.Consume(conf.OrderQueue, conf.OrderConsumer, true, false, false, false, nil)
	if err != nil {
		//记录日志
		fmt.Println(err)
		return
	}

	for item := range result {
		var order model.Order
		err := json.Unmarshal(item.Body, &order)
		if err != nil {
			//记录日志
			fmt.Println(err)
			continue
		}
		if order.Status == model.OrderCreating {
			o := model.Order{}
			if order.Confirmed && order.Price > 0 {
				order.Status = model.SellerConfirmed
				b, err := o.Update(model.SellerConfirmed, order.Address, order.ID, order.AccountId,
					order.ProductId, order.DeliveryId, order.CheckOutId, order.PointId, order.Price)
				if err != nil {
					fmt.Println(err)
					continue
				}
				fmt.Println(b)
				sendOrderMessage2DeliverySystem(order)
			} else {
				b, err := o.Update(model.OrderCreateFailed, order.Address, order.ID, order.AccountId,
					order.ProductId, order.DeliveryId, order.CheckOutId, order.PointId, order.Price)
				if err != nil {
					fmt.Println(err)
					continue
				}
				fmt.Println(b)
			}
		}
		item.Ack(true)
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
