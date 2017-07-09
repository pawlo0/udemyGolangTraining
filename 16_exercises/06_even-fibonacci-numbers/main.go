package main

import "fmt"

func main() {
	var fib, total, first int
	second := 1
	for fib < 4000000 {
		fib = first + second
		first = second
		second = fib
		if fib%2 == 0 {
			total += fib
		}
	}
	fmt.Println(total)
}
