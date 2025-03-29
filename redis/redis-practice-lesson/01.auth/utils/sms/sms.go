package sms

import (
	"context"
	"errors"
	"fmt"
	"redis-parctice-lesson/global"
	"time"
)

// SendSms 发送短信
func SendSms(mobile, msg string) error {
	if mobile == "" {
		return errors.New("手机号码不能为空")
	}
	//TODO 验证手机格式是否正确
	//TODO 验证发送内容是否正确
	key := fmt.Sprintf("sms:%s", mobile)
	r, err := global.GoRedisClient.Set(context.Background(), key, msg, 300*time.Second).Result()
	if err != nil {
		return err
	}
	fmt.Println(r)
	result := fmt.Sprintf("发送成功,%s,%s", mobile, msg)
	fmt.Println(result)
	return nil
}
