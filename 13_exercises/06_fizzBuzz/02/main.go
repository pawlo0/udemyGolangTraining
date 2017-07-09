package main

import (
	"fmt"
	"strconv"
)

func main() {
	for i := 0; i <= 100; i++ {
		result := strconv.Itoa(i) + " "
		if i%3 == 0 {
			result += "Fizz"
		}
		if i%5 == 0 {
			result += "Buzz"
		}
		fmt.Println(result)
	}
}
