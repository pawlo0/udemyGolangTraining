package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Now()
	go foo()
	go bar()
	fmt.Println(time.Since(start))
}

func foo() {
	for i := 0; i < 1000; i++ {
		fmt.Println("Foo:", i)
	}
}

func bar() {
	for i := 0; i < 1000; i++ {
		fmt.Println("Bar:", i)
	}
}
