package main

import (
	cConsumer "rabbitmq-pratice-lesson/checkout/consumer"
	dConsumer "rabbitmq-pratice-lesson/delivery/consumer"
	"rabbitmq-pratice-lesson/model"
	"rabbitmq-pratice-lesson/order/consumer"
	pConsumer "rabbitmq-pratice-lesson/point/consumer"
	sConsumer "rabbitmq-pratice-lesson/seller/consumer"
)

func init() {
	model.InitDB()

	go func() {
		for {
			consumer.ConsumeOrder()
		}
	}()

	go func() {
		for {
			pConsumer.ConsumePoint()
		}
	}()

	go func() {
		for {
			dConsumer.ConsumeDelivery()
		}
	}()

	go func() {
		for {
			sConsumer.ConsumeSeller()
		}
	}()

	go func() {
		for {
			cConsumer.ConsumeCheckout()
		}
	}()

}

func main() {
	c := make(chan struct{})
	<-c
}
