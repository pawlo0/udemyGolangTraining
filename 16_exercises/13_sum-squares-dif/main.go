package main

import (
	"fmt"
)

func main() {
	numbers := 100
	var sumOfSquares, squareOfSum, sum int
	for i := 1; i <= numbers; i++ {
		sumOfSquares += i * i
		sum += i
	}
	squareOfSum = sum * sum

	fmt.Println(squareOfSum - sumOfSquares)
}
