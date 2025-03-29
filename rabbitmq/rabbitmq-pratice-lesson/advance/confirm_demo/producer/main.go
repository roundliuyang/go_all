package main

import (
	"fmt"
	"github.com/streadway/amqp"
	"rabbitmq-pratice-lesson/lib"
	"strconv"
)

const (
	OneExchange   = "OneExchange"
	OneKey        = "OneKey"
	ManyExchange  = "ManyExchange"
	ManyKey       = "ManyKey"
	AsyncExchange = "AsyncExchange"
	AsyncKey      = "AsyncKey"
	DIRECT        = "direct"
)

// One2One 单条确认
func One2One() error {
	channel, err := lib.GetChannel()
	if err != nil {
		return err
	}
	defer channel.Close()
	err = channel.ExchangeDeclare(OneExchange, DIRECT,
		true, false, false, false, nil)
	if err != nil {
		return err
	}
	chConfirm := make(chan amqp.Confirmation, 1)
	confirms := channel.NotifyPublish(chConfirm)
	err = channel.Confirm(false)
	if err != nil {
		return err
	}

	err = channel.Publish(OneExchange, OneKey, false, false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte("从0到Go语言微服务架构师"),
		})
	if err != nil {
		return err
	} else {
		fmt.Println("send suc.")
	}
	confirmed := <-confirms
	if confirmed.Ack {
		fmt.Println("发送成功:", confirmed.DeliveryTag)
	} else {
		fmt.Println("发送失败:", confirmed.DeliveryTag)
	}
	return nil
}

// ManyConfirm 多条确认
func ManyConfirm() error {
	channel, err := lib.GetChannel()
	if err != nil {
		return err
	}
	defer channel.Close()
	wait := 5
	chConfirm := make(chan amqp.Confirmation, wait)
	confirms := channel.NotifyPublish(chConfirm)
	channel.Confirm(false)
	if err != nil {
		return err
	}
	err = channel.ExchangeDeclare(ManyExchange, DIRECT, true,
		false, false, false, nil)
	if err != nil {
		return err
	}
	for i := 0; i < wait; i++ {
		err := channel.Publish(ManyExchange, ManyKey, false, false, amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte("Go语言微服务架构核心22讲-" + strconv.Itoa(i)),
		})
		if err != nil {
			break
		} else {
			fmt.Println("send->", i)
		}
	}
	for wait != 0 {
		confirmed := <-confirms
		if confirmed.Ack {
			fmt.Println("发送成功:", confirmed.DeliveryTag)
		} else {
			fmt.Println("发送失败:", confirmed.DeliveryTag)
		}
		wait--
	}
	return nil
}

// AsyncConfirm 异步确认
func AsyncConfirm() error {
	channel, err := lib.GetChannel()
	if err != nil {
		return err
	}
	defer channel.Close()
	err = channel.ExchangeDeclare(AsyncExchange, DIRECT, true, false, false, false, nil)
	for i := 0; i < 7; i++ {
		transactionPublish(channel, i)
	}
	return nil
}
func transactionPublish(channel *amqp.Channel, i int) error {
	err := channel.Tx()
	if err != nil {
		return err
	}
	defer func() {
		if err != nil {
			channel.TxRollback()
		} else {
			channel.TxCommit()
		}
	}()
	err = channel.Publish(AsyncExchange, AsyncKey, false, false, amqp.Publishing{
		ContentType: "text/plain",
		Body:        []byte("Go语言极简一本通-" + strconv.Itoa(i)),
	})
	if err != nil {
		return err
	}
	return nil
}

func main() {
	//One2One()
	//ManyConfirm()
	AsyncConfirm()
}
