package main

import (
	"fmt"
)

type MyInterface interface{}

func main() {
	var n *int
	fmt.Println(n == nil)
	test(n)
}

func test(r MyInterface) {
	switch r.(type) {
	case *int:
		fmt.Println("Success!")
	default:
		return
	}
}
