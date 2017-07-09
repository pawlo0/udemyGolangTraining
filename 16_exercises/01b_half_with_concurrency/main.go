package main

import "fmt"
import "time"

func main() {
	start := time.Now()
	var cumm int
	var countEvens int

	c := make(chan int)

	go func() {
		for i := 1; i < 100000000; i++ {
			c <- half(i)
		}
		close(c)
	}()

	for n := range c {
		cumm += n
		if n%2 == 0 {
			countEvens++
		}
	}

	fmt.Println(cumm, countEvens, time.Since(start))
}

func half(n int) int {
	return n / 2
}
