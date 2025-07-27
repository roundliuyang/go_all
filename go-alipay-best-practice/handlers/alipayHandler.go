package handlers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-pay/gopay/alipay"
	"go-alipay-best-practice/config"
	"go.uber.org/zap"
	"net/http"
)

func AlipayNotify(c *gin.Context) {
	tradeStatus := c.PostForm("trade_status")
	if tradeStatus == "TRADE_CLOSED" {
		c.JSON(http.StatusOK, gin.H{
			"msg": "交易已关闭",
		})
	}
	fmt.Println("成功-----------------------------------------------------")
	if tradeStatus == "TRADE_SUCCESS" {
		// 验签
		// TODO 做自己的业务，订单状态的修改/安排物流/...
		c.JSON(http.StatusOK, gin.H{
			"msg": "成功",
		})
	}

}

func AlipayReturn(c *gin.Context) {
	log, _ := zap.NewProduction()
	notifyReq, err := alipay.ParseNotifyToBodyMap(c.Request)
	if err != nil {
		log.Error("alipay.ParseNotifyToBodyMap", zap.Error(err))
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": "参数错误",
		})
		return
	}
	ok, err := alipay.VerifySign(config.AliPublicKey, notifyReq)
	if err != nil {
		log.Error("alipay.VerifySign", zap.Error(err))
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": "参数错误",
		})
	}
	msg := ""
	if ok {
		msg = "验签成功"
	} else {
		msg = "验签失败"
	}
	// TODO 做自己的业务，订单状态的修改/安排物流/...
	c.JSON(http.StatusOK, gin.H{
		"msg": msg,
	})
}
