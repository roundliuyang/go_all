package main

import (
	"fmt"
	"math/rand"
	"rabbitmq-pratice-lesson/advance/dl_demo/internal"
	"rabbitmq-pratice-lesson/lib"
)

func normalConsumer() {
	channel, err := lib.GetChannel()
	if err != nil {
		fmt.Println(err)
		return
	}
	defer channel.Close()
	msgList, err := channel.Consume(internal.NormalQueue, internal.NormalRoutingKey, false, false, false, false, nil)
	if err != nil {
		fmt.Println(err)
		return
	}
	for d := range msgList {
		data := rand.Int63n(100)
		fmt.Println(data)
		if data%2 == 0 {
			fmt.Printf("Normal...%s", string(d.Body))
			d.Ack(true)
		} else {
			d.Reject(false)
		}

	}
}

func main() {
	c := make(chan struct{})
	go func() {
		normalConsumer()
	}()
	<-c
}
