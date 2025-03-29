package lib

import (
	"fmt"
	"github.com/streadway/amqp"
	"rabbitmq-pratice-lesson/global"
)

func GetMQURL(userName, password, ip string, port int) string {
	mqUrl := fmt.Sprintf("amqp://%s:%s@%s:%d", userName, password, ip, port)
	return mqUrl
}

func GetChannel() (*amqp.Channel, error) {
	url := GetMQURL("root", "123456", global.HOST, global.PORT)
	conn, err := amqp.Dial(url)
	if err != nil {
		//TODO 生产环境，你要写日志
		panic(err)
	}
	//channel, err := conn.Channel()
	//return channel, err

	return conn.Channel()
}
