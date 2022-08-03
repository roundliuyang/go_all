package unknown

import (
	"fmt"
	"testing"
)

type Lesson struct {
	name, target string
	spend        int
}

func TestFibList(t *testing.T) {

	lesson8 := &Lesson{"从0到Go语言微服务架构师", "全面掌握Go语言微服务如何落地，代码级彻底一次性解决微服务和分布式系统。", 50}

	lesson8.name = "1"
	fmt.Println("lesson8 name: ", (*lesson8).name)
	fmt.Println("lesson8 name: ", lesson8.name)
}
