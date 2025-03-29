package main

import (
	"fmt"
	"rabbitmq-pratice-lesson/conf"
	"rabbitmq-pratice-lesson/lib"
)

const (
	OneExchange   = "OneExchange"
	OneQueue      = "OneQueue"
	OneKey        = "OneKey"
	OneConsumer   = "OneConsumer"
	ManyExchange  = "ManyExchange"
	ManyQueue     = "ManyQueue"
	ManyConsumer  = "ManyConsumer"
	ManyKey       = "ManyKey"
	AsyncExchange = "AsyncExchange"
	AsyncQueue    = "AsyncQueue"
	AsyncConsumer = "AsyncConsumer"
	AsyncKey      = "AsyncKey"
)

func Consume(exchange, q, consumer, key, mode string) error {
	channel, err := lib.GetChannel()
	if err != nil {
		return err
	}

	err = channel.ExchangeDeclare(exchange, mode,
		true, false, false, false, nil)
	if err != nil {
		//记录日志
		fmt.Println(err)
		return err
	}
	_, err = channel.QueueDeclare(q,
		true, false, false, false, nil)
	if err != nil {
		//记录日志
		fmt.Println(err)
		return err
	}

	err = channel.QueueBind(q, key, exchange,
		false, nil)
	if err != nil {
		//记录日志
		fmt.Println(err)
		return err
	}

	result, err := channel.Consume(q, consumer, false,
		false, false, false, nil)
	if err != nil {
		return err
	}
	for item := range result {
		fmt.Println(string(item.Body))
		item.Ack(true)
	}
	return nil
}

func main() {
	c := make(chan struct{})
	go func() {
		Consume(OneExchange, OneQueue, OneConsumer, OneKey, conf.DIRECT)
	}()

	go func() {
		Consume(ManyExchange, ManyQueue, ManyConsumer, ManyKey, conf.DIRECT)
	}()
	go func() {
		Consume(AsyncExchange, AsyncQueue, AsyncConsumer, AsyncKey, conf.DIRECT)
	}()

	<-c

}
