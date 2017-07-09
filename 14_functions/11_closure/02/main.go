package main

import (
	"fmt"
)

var x int

func main() {
	fmt.Println(increment())
	fmt.Println(increment())
}

func increment() int {
	x++
	return x
}
