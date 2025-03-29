package main

import (
	"fmt"
	"github.com/streadway/amqp"
	"rabbitmq-pratice-lesson/lib"
)

const (
	exchange   = "exchange.return"
	exMode     = "direct"
	routingKey = "routeReturn"
)

func main() {
	channel, err := lib.GetChannel()
	if err != nil {
		fmt.Println(err)
		return
	}
	defer channel.Close()

	err = channel.ExchangeDeclare(exchange, exMode,
		true, false, false, false, nil)
	if err != nil {
		fmt.Println(err)
		return
	}
	chReturn := make(chan amqp.Return)
	go func() {
		notifyReturn := channel.NotifyReturn(chReturn)
		for r := range notifyReturn {
			fmt.Printf("路由失败原因:%s,返回消息:%s", r.ReplyText, string(r.Body))
		}
	}()
	err = channel.Publish(exchange, routingKey, true, false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte("从0到Go语言微服务架构师"),
		})
	if err != nil {
		fmt.Println(err)
	}

}
