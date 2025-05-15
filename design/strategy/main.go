package main

type PaymentStrategy interface {
	TransactionFee(amount float64) float64
}

type CreditCard struct{}

func (c *CreditCard) TransactionFee(amount float64) float64 {
	return 0.03 * amount // 假设 3% 的手续费
}

type PayPal struct{}

func (p *PayPal) TransactionFee(amount float64) float64 {
	return 0.04 * amount // 假设 4% 的手续费
}

type Bitcoin struct{}

func (b *Bitcoin) TransactionFee(amount float64) float64 {
	return 0.01 * amount // 假设 1% 的手续费
}

// -------------------------------------------------------------------------------------------------------
type Context struct {
	strategy PaymentStrategy
}

func (c *Context) SetStrategy(strategy PaymentStrategy) {
	c.strategy = strategy
}

func (c *Context) TransactionFee(amount float64) {
	c.strategy.TransactionFee(amount)
}

func NewContext() *Context {
	return &Context{}
}

// -----------------------------------------------------------------------------------------------------------

func main() {
	strategies := map[string]PaymentStrategy{
		"CreditCard": &CreditCard{},
		"PayPal":     &PayPal{},
		"Bitcoin":    &Bitcoin{},
	}

	userChoice := "CreditCard"
	strategy := strategies[userChoice]
	strategy.TransactionFee(4.1)

	// 或者
	context := NewContext()
	// 注意， 只有指针实现了接口
	card := &CreditCard{}
	context.SetStrategy(card)

	context.TransactionFee(10.1)
}
