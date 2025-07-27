package config

const (
	AppId = "9021000150676293"
	// 应用私钥
	PrivateKey = ""
	// 支付宝公钥
	AliPublicKey   = ""
	IsProduction   = false
	ProductionCode = "FAST_INSTANT_TRADE_PAY"
	Host           = "todo.free.idcfengye.com"
	NotifyURL      = "http://" + Host + "/pay/alipay/notify"
	ReturnURL      = "http://127.0.0.1:8080/pay/alipay/return"
)
