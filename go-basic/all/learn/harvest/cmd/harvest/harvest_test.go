package main

import (
	"fmt"
	"net/url"
	"testing"
)

func TestUnescape(t *testing.T) {
	v := url.Values{}
	v.Add("cat sounds", "meow")
	v.Add("cat sounds", "mew")
	fmt.Println(v.Get("cat sounds"))
}
