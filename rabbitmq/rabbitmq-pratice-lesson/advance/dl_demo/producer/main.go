package main

import (
	"fmt"
	"github.com/streadway/amqp"
	"rabbitmq-pratice-lesson/advance/dl_demo/internal"
	"rabbitmq-pratice-lesson/lib"
	"strconv"
	"time"
)

func main() {
	channel, err := lib.GetChannel()
	if err != nil {
		fmt.Println(err)
		return
	}
	err = channel.ExchangeDeclare(internal.NormalExchange, amqp.ExchangeDirect,
		true, false, false, false, nil)
	if err != nil {
		fmt.Println(err)
		return
	}
	_, err = channel.QueueDeclare(internal.NormalQueue,
		true, false, false, false, amqp.Table{
			"x-message-ttl":             60000, // 消息过期时间（队列级别）,毫秒
			"x-dead-letter-exchange":    internal.DeadExchange,
			"x-dead-letter-routing-key": internal.DeadRoutingKey,
		})
	if err != nil {
		fmt.Println(err)
		return
	}
	err = channel.QueueBind(internal.NormalQueue, internal.NormalRoutingKey,
		internal.NormalExchange, false, nil)
	if err != nil {
		fmt.Println(err)
		return
	}

	err = channel.ExchangeDeclare(internal.DeadExchange, amqp.ExchangeFanout,
		true, false, false, false, nil)
	if err != nil {
		fmt.Println(err)
		return
	}

	_, err = channel.QueueDeclare(internal.DeadQueue,
		true, false, false, false, nil)
	if err != nil {
		fmt.Println(err)
		return
	}

	err = channel.QueueBind(internal.DeadQueue, "",
		internal.DeadExchange, false, nil)
	if err != nil {
		fmt.Println(err)
		return
	}

	message := "面向加薪学习-《从0到Go语言微服务架构师》训练营" + strconv.Itoa(int(time.Now().Unix()))
	fmt.Println(message)
	err = channel.Publish(internal.NormalExchange, internal.NormalRoutingKey, false, false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(message),
			Expiration:  "10000", // 消息过期时间,毫秒
		})
	if err != nil {
		fmt.Println(err)
	}
}
