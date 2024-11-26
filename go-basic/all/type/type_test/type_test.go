package type_test

import "testing"

type MyInt int64

func TestImplicit(t *testing.T) {
	var a int32 = 1
	var b int64
	// 显示类型转换，隐式不被允许
	b = int64(a)
	var c MyInt
	c = MyInt(b)
	t.Log(a, b, c)
}

// 不支持指针运算
func TestPoint(t *testing.T) {
	a := 1
	aPtr := &a
	//aPtr = aPtr + 1
	t.Log(a, aPtr)
	t.Logf("%T %T", a, aPtr)
}

// 默认 空值，字符串是值类型
func TestString(t *testing.T) {
	var s string
	t.Log("*" + s + "*") //初始化零值是“”
	t.Log(len(s))

}
