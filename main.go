package main

import (
	"fmt"

	"github.com/ashupednekar/testgopy/pycall"
)

func foo() string {
	fmt.Println("in foo")
	return "a"
}

func main() {
	pc := pycall.NewCaller(10)
	pc.Start(foo)

	for r := range pc.Ch {
		fmt.Println("r: ", r)
	}
}
