package main

import (
	"fmt"
	"time"
)

func main() {

	start := time.Now()

	triangle, i, j, k, divs := 0, 1, 1, 1, 0

	for divs < 498 {

		triangle += i
		divs = 0

		if triangle%2 == 0 {
			j, k = 2, 1
		} else {
			j, k = 3, 2
		}

		for j*j <= triangle {
			if triangle%j == 0 {
				divs += 2
			}
			j += k
		}
		i++

	}
	fmt.Println(triangle, " ", time.Since(start))
}
