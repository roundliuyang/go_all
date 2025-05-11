package interface_test

import (
	"fmt"
	"testing"
)

type Phone interface {
	call()
}

type iPhone struct {
	name string
}

func (phone iPhone) call() {
	fmt.Println("Hello, iPhone.")
}

func (phone iPhone) send_wechat() {
	fmt.Println("Hello, Wechat.")
}

func TestInterface(t *testing.T) {
	var phone Phone
	phone = iPhone{name: "ming's iphone"}
	phone.call()
	//phone.send_wechat()

	phone2 := iPhone{name: "ming's iphone"}
	phone2.call()
	phone2.send_wechat()
}

func printType(i interface{}) {

	switch i.(type) {
	case int:
		fmt.Println("参数的类型是 int")
	case string:
		fmt.Println("参数的类型是 string")
	}
}

func TestInterface2(t *testing.T) {
	a := 10
	printType(a)
}

func TestInterface3(t *testing.T) {
	//a := 10
	//
	//switch a.(type) {
	//case int:
	//	fmt.Println("参数的类型是 int")
	//case string:
	//	fmt.Println("参数的类型是 string")
	//}
}

func TestInterface4(t *testing.T) {
	a := 10

	switch interface{}(a).(type) {
	case int:
		fmt.Println("参数的类型是 int")
	case string:
		fmt.Println("参数的类型是 string")
	}
}

func TestInterface5(t *testing.T) {
	//var a interface{} = 10
	//
	//switch b := a.(type) {
	//case int:
	//	_ = b.(int)
	//	fmt.Println("参数的类型是 int")
	//case string:
	//	fmt.Println("参数的类型是 string")
	//}
}
