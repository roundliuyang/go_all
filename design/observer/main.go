package main

import "fmt"

// Observer 接口：所有关注股票价格的分析师都要实现这个接口
type Observer interface {
	Notify(stockPrice float32)
}

// StockMarket 被观察者，在股票价格变动时，会通知所有注册的分析师
type StockMarket struct {
	observers  []Observer
	stockPrice float32
}

func (s *StockMarket) Register(observer Observer) {
	s.observers = append(s.observers, observer)
}

func (s *StockMarket) SetPrice(price float32) {
	s.stockPrice = price
	s.notifyAll()
}

func (s *StockMarket) notifyAll() {
	for _, observer := range s.observers {
		observer.Notify(s.stockPrice)
	}
}

// Analyst 实现了 Observer 接口，代表一个股票市场的分析师
type Analyst struct {
	name string
}

func (a *Analyst) Notify(stockPrice float32) {
	fmt.Printf("%s 得到了新的股价：$%f\n", a.name, stockPrice)
}

func main() {
	stockMarket := &StockMarket{}

	alice := &Analyst{name: "Alice"}
	bob := &Analyst{name: "Bob"}

	stockMarket.Register(alice)
	stockMarket.Register(bob)

	stockMarket.SetPrice(120.0) // 当价格改变时，Alice 和 Bob 都会被通知
	stockMarket.SetPrice(125.5) // 再次改变价格，Alice 和 Bob 都会得到通知

}
