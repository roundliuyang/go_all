package main

import "fmt"

type Animal interface {
	Speak() string
}

type Dog struct{}

func (d *Dog) Speak() string {
	return "Woof!"
}

type Cat struct{}

func (c *Cat) Speak() string {
	return "Meow!"
}

type PaymentMethod interface {
	Pay(amount float64) string
}

type Paypal struct{}

func (p *Paypal) Pay(amount float64) string {
	return fmt.Sprintf("Paid %f via Paypal", amount)
}

type CreditCard struct{}

func (c *CreditCard) Pay(amount float64) string {
	return fmt.Sprintf("Paid %f using Credit Card", amount)
}

// Animal Map
var AnimalFactory = map[string]func() Animal{
	"dog": func() Animal { return &Dog{} },
	"cat": func() Animal { return &Cat{} },
}

// Payment Map
var PaymentFactory = map[string]func() PaymentMethod{
	"paypal":     func() PaymentMethod { return &Paypal{} },
	"creditcard": func() PaymentMethod { return &CreditCard{} },
}

// Animal Factory
func GetAnimal(t string) Animal {
	if factory, ok := AnimalFactory[t]; ok {
		return factory()
	}
	return nil
}

// Payment Factory
func GetPaymentMethod(method string) PaymentMethod {
	if factory, ok := PaymentFactory[method]; ok {
		return factory()
	}
	return nil
}

func main() {
	animal := GetAnimal("dog")
	fmt.Println(animal.Speak()) // Woof!

	payment := GetPaymentMethod("paypal")
	fmt.Println(payment.Pay(100.5)) // Paid 100.5 via Paypal
}
