package time

import (
	"fmt"
	"testing"
	"time"
)

func TestFibList(t *testing.T) {
	// 2022-08-01 23:23:41.638473 +0800 CST m=+0.011448601
	fmt.Println(time.Now())

	// 返回将 t 向下舍入为 d 的倍数
	fmt.Println(time.Now().Truncate(10 * 48 * time.Hour))

	// 向下舍入为 d 的倍数
	fmt.Println(time.Now().Add(48 * time.Hour))

	// 20220802
	fmt.Println(time.Now().Format("20060102"))

	fmt.Println(time.Now().Add(-30 * 24 * time.Hour))
}
