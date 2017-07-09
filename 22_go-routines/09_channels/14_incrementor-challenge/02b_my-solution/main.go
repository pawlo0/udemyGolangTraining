package main

import (
	"fmt"
	"strconv"
)

func main() {

	c := incrementor(2)

	var count int
	for n := range c {
		count++
		fmt.Println(n)
	}

	fmt.Println("Final Count:", count)
}

func incrementor(n int) chan string {
	c := make(chan string)
	done := make(chan bool)

	output := func(i int) {
		for k := 0; k < 20; k++ {
			c <- fmt.Sprint("Process: "+strconv.Itoa(i)+" printing:", k)
		}
		done <- true
	}

	for i := 1; i < n+1; i++ {
		go output(i)
	}

	go func() {
		for i := 0; i < n; i++ {
			<-done
		}
		close(c)
	}()

	return c
}
