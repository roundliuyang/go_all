package sms

import (
	"redis-parctice-lesson/global"
	"testing"
)

func init() {
	global.InitGoRedisClient()
}

func TestSendSms(t *testing.T) {
	mobile := "130000000"
	msg := "123456"
	err := SendSms(mobile, msg)
	if err != nil {
		panic(err)
	}
}
