package main

import (
	"fmt"
)

func main() {
	c := make(chan int)
	go func() {
		c <- 1
	}()
	fmt.Println(<-c)
}

// This results in a deadlock
// chan you determine why?
// And what would you do to fix it?
