package main

import (
	"fmt"
	"rabbitmq-pratice-lesson/advance/dl_demo/internal"
	"rabbitmq-pratice-lesson/lib"
)

func dlConsumer() {
	channel, err := lib.GetChannel()
	if err != nil {
		fmt.Println(err)
		return
	}
	defer channel.Close()
	msgList, err := channel.Consume(internal.DeadQueue, internal.DeadRoutingKey,
		false, false, false, false, nil)
	if err != nil {
		fmt.Println(err)
		return
	}
	for d := range msgList {
		fmt.Printf("DL...%s\n", string(d.Body))
		d.Ack(true)
	}
}

func main() {
	c := make(chan struct{})
	go func() {
		dlConsumer()
	}()
	<-c
}
