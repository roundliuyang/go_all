package strconv

import (
	"fmt"
	"strconv"
	"testing"
)

func TestFibList(t *testing.T) {

	// 字符串 到 int
	fmt.Println(strconv.Atoi("-42"))

	// int 到 字符串
	fmt.Println(strconv.Itoa(-42))
}
