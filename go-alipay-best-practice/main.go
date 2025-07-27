package main

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-pay/gopay"
	"github.com/go-pay/gopay/alipay"
	"go-alipay-best-practice/config"
	"go-alipay-best-practice/handlers"
	"go.uber.org/zap"
	"time"
)

var log *zap.Logger

func main() {
	log, _ := zap.NewProduction()
	// 获取 url 进行支付
	client, err := alipay.NewClient(config.AppId, config.PrivateKey, config.IsProduction)
	if err != nil {
		log.Error("NewClient err", zap.Error(err))
		return
	}
	client.SetCharset("utf-8").
		SetSignType(alipay.RSA2).
		SetNotifyUrl(config.NotifyURL).
		SetReturnUrl(config.ReturnURL)

	ts := time.Now().UnixMilli()
	fmt.Println("OutTradeNo:", ts)

	outTradeNo := fmt.Sprintf("%d", ts)
	bm := make(gopay.BodyMap)
	bm.Set("subject", "测试支付")
	bm.Set("out_trade_no", outTradeNo)
	bm.Set("total_amount", "88.88")
	bm.Set("product_code", config.ProductionCode)

	payUrl, err := client.TradePagePay(context.Background(), bm)
	if err != nil {
		log.Error("TradePagePay err", zap.Error(err))
		return
	}
	log.Info("payUrl", zap.String("payUrl", payUrl))

	// 支付成功后，支付宝回调我们
	engine := gin.Default()
	engine.POST("/pay/alipay/notify", handlers.AlipayNotify)
	engine.GET("/pay/alipay/return", handlers.AlipayReturn)
	engine.Run(":8080")

}
