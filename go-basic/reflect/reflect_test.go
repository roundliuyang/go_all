package reflect_test

import (
	"fmt"
	"reflect"
	"testing"
)

func TestTypeAndValue(t *testing.T) {
	// 第一定律
	var age interface{} = 25

	fmt.Printf("原始接口变量的类型为 %T，值为 %v \n", age, age)

	tt := reflect.TypeOf(age)
	v := reflect.ValueOf(age)

	// 从接口变量到反射对象
	fmt.Printf("从接口变量到反射对象：Type对象的类型为 %T \n", tt)
	fmt.Printf("从接口变量到反射对象：Value对象的类型为 %T \n", v)

	// 第二定律
	// 从反射对象到接口变量
	i := v.Interface()
	fmt.Printf("从反射对象到接口变量：新对象的类型为 %T 值为 %v \n", i, i)

	// 当然了，最后转换后的对象，静态类型为 interface{} ，如果要转成最初的原始类型，需要再类型断言转换一下
	i = v.Interface().(int)
	fmt.Printf("从反射对象到接口变量：新对象的类型为 %T 值为 %v \n", i, i)
}

func TestSettable(t *testing.T) {
	//var name string = "Go编程时光"
	//
	//v := reflect.ValueOf(name)
	//fmt.Println("可写性为:", v.CanSet())

	var name string = "Go编程时光"
	v1 := reflect.ValueOf(&name)
	fmt.Println("v1 可写性为:", v1.CanSet())

	v2 := v1.Elem()
	fmt.Println("v2 可写性为:", v2.CanSet())
}

func TestSettable2(t *testing.T) {
	var name string = "Go编程时光"
	fmt.Println("真实世界里 name 的原始值为：", name)

	v1 := reflect.ValueOf(&name)
	v2 := v1.Elem()

	v2.SetString("Python编程时光")
	fmt.Println("通过反射对象进行更新后，真实世界里 name 变为：", name)
}

func CheckType(v interface{}) {
	t := reflect.TypeOf(v)
	switch t.Kind() {
	case reflect.Float32, reflect.Float64:
		fmt.Println("Float")
	case reflect.Int, reflect.Int32, reflect.Int64:
		fmt.Println("Integer")
	default:
		fmt.Println("Unknown", t)
	}
}

func TestBasicType(t *testing.T) {
	var f float64 = 12
	CheckType(f)
	CheckType(&f)

}

func TestDeepEqual(t *testing.T) {
	a := map[int]string{1: "one", 2: "two", 3: "three"}
	b := map[int]string{1: "one", 2: "two", 3: "three"}
	//	t.Log(a == b)
	t.Log("a==b?", reflect.DeepEqual(a, b))

	s1 := []int{1, 2, 3}
	s2 := []int{1, 2, 3}
	s3 := []int{2, 3, 1}

	t.Log("s1 == s2?", reflect.DeepEqual(s1, s2))
	t.Log("s1 == s3?", reflect.DeepEqual(s1, s3))

	c1 := Customer{"1", "Mike", 40}
	c2 := Customer{"1", "Mike", 40}
	fmt.Println(c1 == c2)
	fmt.Println(reflect.DeepEqual(c1, c2))
}

type Employee struct {
	EmployeeID string
	Name       string `format:"normal"`
	Age        int
}

func (e *Employee) UpdateAge(newVal int) {
	e.Age = newVal
}

type Customer struct {
	CookieID string
	Name     string
	Age      int
}

func TestInvokeByName(t *testing.T) {
	e := &Employee{"1", "Mike", 30}
	//按名字获取成员
	t.Logf("Name: value(%[1]v), Type(%[1]T) ", reflect.ValueOf(*e).FieldByName("Name"))
	if nameField, ok := reflect.TypeOf(*e).FieldByName("Name"); !ok {
		t.Error("Failed to get 'Name' field.")
	} else {
		t.Log("Tag:format", nameField.Tag.Get("format"))
	}
	reflect.ValueOf(e).MethodByName("UpdateAge").
		Call([]reflect.Value{reflect.ValueOf(1)})
	t.Log("Updated Age:", e)
}
