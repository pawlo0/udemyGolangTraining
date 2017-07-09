package main

import "fmt"
import "time"

func main() {
	start := time.Now()
	var cumm int
	var countEvens int
	for i := 1; i < 100000000; i++ {
		value, isEven := half(i)
		cumm += value
		if isEven {
			countEvens++
		}
	}
	fmt.Println(cumm, countEvens, time.Since(start))
}

func half(n int) (int, bool) {
	return n / 2, n%2 == 0
}
