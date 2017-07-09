package main

import (
	"fmt"
)

func startsAt(x int, y int) func() int {
	return func() int {
		x += y
		return x
	}
}

func main() {
	result1 := startsAt(5, 1)
	result2 := startsAt(10, 10)
	fmt.Println(result1())
	fmt.Println(result1())
	fmt.Println(result1())
	fmt.Println(result2())
	fmt.Println(result2())
}
