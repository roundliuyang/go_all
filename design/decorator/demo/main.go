package main

import "fmt"

type Beverage interface {
	Description() string
	Cost() float64
}

type Coffee struct{}

func (c *Coffee) Description() string {
	return "基本咖啡"
}

func (c *Coffee) Cost() float64 {
	return 5.0
}

type MilkDecorator struct {
	beverage Beverage
}

func (m *MilkDecorator) Description() string {
	return m.beverage.Description() + " + 牛奶"
}

func (m *MilkDecorator) Cost() float64 {
	return m.beverage.Cost() + 1.5
}

type ExtraFeature func(Beverage) Beverage

func WithCaramel(b Beverage) Beverage {
	return &CaramelDecorator{beverage: b}
}

type CaramelDecorator struct {
	beverage Beverage
}

func (c *CaramelDecorator) Description() string {
	return c.beverage.Description() + " + 焦糖"
}

func (c *CaramelDecorator) Cost() float64 {
	return c.beverage.Cost() + 2.0
}

func main() {
	coffee := &Coffee{}
	fmt.Println(coffee.Description(), "$", coffee.Cost())

	// 使用裝飾者加牛奶
	milkCoffee := &MilkDecorator{beverage: coffee}
	fmt.Println(milkCoffee.Description(), "$", milkCoffee.Cost())

	// 使用回調函數加焦糖
	caramelCoffee := WithCaramel(coffee)
	fmt.Println(caramelCoffee.Description(), "$", caramelCoffee.Cost())
}
