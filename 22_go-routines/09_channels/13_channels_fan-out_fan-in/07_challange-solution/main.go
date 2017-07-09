package main

import (
	"fmt"
	"sync"
)

func main() {

	in := gen()

	// FAN OUT
	f1 := factorial(in)
	f2 := factorial(in)
	f3 := factorial(in)
	f4 := factorial(in)
	f5 := factorial(in)
	f6 := factorial(in)
	f7 := factorial(in)
	f8 := factorial(in)
	f9 := factorial(in)
	f0 := factorial(in)

	// FAN IN
	var y int
	for n := range merge(f1, f2, f3, f4, f5, f6, f7, f8, f9, f0) {
		y++
		fmt.Println(y, "\t", n)
	}
}

func gen() <-chan int {
	out := make(chan int)
	go func() {
		for i := 0; i < 100; i++ {
			for j := 3; j < 13; j++ {
				out <- j
			}
		}
		close(out)
	}()
	return out
}

func factorial(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		for n := range in {
			out <- fact(n)
		}
		close(out)
	}()
	return out
}

func fact(n int) int {
	total := 1
	for i := n; i > 0; i-- {
		total *= i
	}
	return total
}

func merge(cs ...<-chan int) <-chan int {
	out := make(chan int)
	var wg sync.WaitGroup

	output := func(c <-chan int) {
		for n := range c {
			out <- n
		}
		wg.Done()
	}

	wg.Add(len(cs))

	for _, c := range cs {
		go output(c)
	}

	// Start a go routine to close out once all output goroutines are done.
	// This must start after the wg.Add call
	go func() {
		wg.Wait()
		close(out)
	}()

	return out
}

/*
CHALLENGE #1:
-- Change the code above to execute 1,000 factorial computations concurrently and in parallel.
-- Use the "fan out / fan in" pattern
CHALLENGE #2:
WATCH MY SOLUTION BEFORE DOING THIS CHALLENGE #2
-- While running the factorial computations, try to find how much of your resources are being used.
-- Post the percentage of your resources being used to this discussion: https://goo.gl/BxKnOL
*/
