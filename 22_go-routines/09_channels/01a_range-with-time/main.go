package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Now()
	c := make(chan int)

	go func() {
		for i := 0; i < 100000; i++ {
			c <- i
		}
		close(c)
	}()

	for n := range c {
		fmt.Println(n)
	}

	fmt.Println(time.Since(start))

}
