package consumer

import (
	"encoding/json"
	"fmt"
	"github.com/streadway/amqp"
	"math/rand"
	"rabbitmq-pratice-lesson/conf"
	"rabbitmq-pratice-lesson/lib"
	"rabbitmq-pratice-lesson/model"
	"time"
)

func ConsumeCheckout() {
	channel, err := lib.GetChannel()
	if err != nil {
		//记录日志
		fmt.Println(err)
		return
	}
	err = channel.ExchangeDeclare(conf.CheckoutExchange, conf.DIRECT,
		true, false, false, false, nil)
	if err != nil {
		//记录日志
		fmt.Println(err)
		return
	}
	_, err = channel.QueueDeclare(conf.CheckoutQueue, true,
		false, false, false, nil)
	if err != nil {
		//记录日志
		fmt.Println(err)
		return
	}

	err = channel.QueueBind(conf.CheckoutQueue, conf.CheckoutKey, conf.CheckoutExchange, false, nil)
	if err != nil {
		//记录日志
		fmt.Println(err)
		return
	}
	result, err := channel.Consume(conf.CheckoutQueue, conf.CheckoutConsumer,
		true, false, false, false, nil)
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
		p := model.Point{}
		point, err := p.Add(order.ID, order.Price, conf.CheckoutSUCCESS)
		if err != nil {
			//记录日志
			fmt.Println(err)
			continue
		}
		order.PointId = point.ID
		rand.Seed(time.Now().UnixNano())
		randNum := rand.Intn(10000)
		s := model.CheckOut{}
		result, err := s.Save(order.ID, int64(randNum), order.Price, conf.CheckoutSUCCESS)
		if err != nil {
			//记录日志
			fmt.Println(err)
			continue
		}
		order.CheckOutId = result.ID
		checkOutProcessOrder(order)
		item.Ack(true)

	}
}
func checkOutProcessOrder(order model.Order) {
	o := model.Order{}
	if order.Status == model.DeliveryConformed && order.CheckOutId > 0 {
		order.Status = model.CheckoutConformed
		b, err := o.Update(model.CheckoutConformed, order.Address, order.ID, order.AccountId,
			order.ProductId, order.DeliveryId, order.CheckOutId, order.PointId, order.Price)
		if err != nil {
			//单体服务，如何处理
			//微服务，如何处理
		}
		fmt.Println(b)
		sendOrderMessage2PointSystem(order)
	}
}

func sendOrderMessage2PointSystem(o model.Order) {
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
	err = channel.Publish(conf.PointExchange, conf.PointRoutingKey, false, false, msg)
	if err != nil {
		//记录日志
		fmt.Println(err)
		return
	}
}
