package main

import "fmt"

func oddOrEven(x int) string {
	remainder := x % 2

	if remainder == 1 {
		return "Odd"
	}
	return "Even"
}

func main() {
	for i := 0; i < 100; i++ {
		fmt.Println(i, oddOrEven(i))
	}
}
