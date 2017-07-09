package main

import (
	"fmt"
)

func main() {
	customenrNumber := make([]int, 3)
	// 3 is length and capacity

	customenrNumber[0] = 7
	customenrNumber[1] = 10
	customenrNumber[2] = 15

	fmt.Println(customenrNumber[0])
	fmt.Println(customenrNumber[1])
	fmt.Println(customenrNumber[2])

	greeting := make([]string, 3, 5)
	// 3 is length - number of elements referred to by the slice
	// 5 is capacity - number of elements in the underlying array

	greeting[0] = "Good morning!"
	greeting[1] = "Bonjour!"
	greeting[2] = "dias!"

	fmt.Println(greeting[2])
}
