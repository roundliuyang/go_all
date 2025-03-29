package conf

const (
	DIRECT  = "direct"
	FANOUT  = "fanout"
	TOPIC   = "topic"
	HEADERS = "headers"
)

const (
	CheckoutSUCCESS = iota + 1
	CheckoutFAILED
)

const (
	PointSUCCESS = iota + 1
	PointFAILED
)

const (
	ProductUp = iota + 1
	ProductDown
)

const (
	SellerOPEN = iota + 1
	SellerCLOSE
)

const (
	DeliveryExchange = "exchange.delivery"
	DeliveryQueue    = "queue.delivery"
	DeliveryKey      = "key.delivery"

	SellerExchange = "exchange.seller"
	SellerQueue    = "queue.seller"
	SellerKey      = "key.seller"

	CheckoutExchange = "exchange.checkout"
	CheckoutQueue    = "queue.checkout"
	CheckoutKey      = "key.checkout"

	PointExchange   = "exchange.point"
	PointQueue      = "queue.point"
	PointRoutingKey = "key.point.topic"
	PointConsumeKey = "#.point.topic"

	OrderExchange   = "exchange.order"
	OrderQueue      = "queue.order"
	OrderRoutingKey = "key.order.topic"
	OrderConsumeKey = "#.order.topic"

	OrderConsumer    = "consumer.order"
	DeliveryConsumer = "consumer.delivery"
	CheckoutConsumer = "consumer.checkout"
	PointConsumer    = "consumer.point"
	SellerConsumer   = "consumer.seller"
)
