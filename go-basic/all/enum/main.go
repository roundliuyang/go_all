package main

import (
	"all/all/enum/policy_type"
	"fmt"
)

func main() {
	foo(policy_type.Policy_MAX)
}

func foo(p policy_type.PolicyType) {
	fmt.Printf("enum value: %v\n", p)
}
