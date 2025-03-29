package consumer

import (
	"encoding/json"
	"fmt"
	"rabbitmq-pratice-lesson/conf"
	"rabbitmq-pratice-lesson/lib"
	"rabbitmq-pratice-lesson/model"
)

func ConsumePoint() {
	channel, err := lib.GetChannel()
	if err != nil {
		//记录日志
		fmt.Println(err)
		return
	}
	err = channel.ExchangeDeclare(conf.PointExchange, conf.TOPIC, true, false,
		false, false, nil)
	if err != nil {
		//记录日志
		fmt.Println(err)
		return
	}
	_, err = channel.QueueDeclare(conf.PointQueue, true, false, false, false, nil)
	if err != nil {
		//记录日志
		fmt.Println(err)
		return
	}
	err = channel.QueueBind(conf.PointQueue, conf.PointConsumeKey, conf.PointExchange, false, nil)
	if err != nil {
		//记录日志
		fmt.Println(err)
		return
	}
	result, err := channel.Consume(conf.PointQueue, conf.PointConsumer, true,
		false, false, false, nil)
	if err != nil {
		//记录日志
		fmt.Println(err)
		return
	}
	for item := range result {
		var order model.Order
		err := json.Unmarshal(item.Body, &order)
		if err != nil {
			//记录日志，设置报警机制
			fmt.Println(err)
			continue
		}
		p := model.Point{}
		point, err := p.Add(order.ID, order.Price, conf.PointSUCCESS)
		if err != nil {
			//记录日志，设置报警机制
			fmt.Println(err)
			continue
		}
		order.PointId = point.ID
		pointProcessOrder(order)
		item.Ack(true)
	}
}
func pointProcessOrder(order model.Order) {
	o := model.Order{}
	if order.Status == model.CheckoutConformed && order.PointId > 0 {
		b, err := o.Update(model.OrderCreated, order.Address, order.ID, order.AccountId, order.ProductId, order.DeliveryId,
			order.CheckOutId, order.PointId, order.Price)
		if err != nil {
			if err != nil {
				//记录日志
				fmt.Println(err)
				return
			}
		}
		fmt.Println(b)
	}
}
