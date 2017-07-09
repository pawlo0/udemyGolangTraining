package main

import (
	"fmt"
)

func main() {
	greet("John")
	greet("Jane")
}

func greet(name string) {
	fmt.Println("Hello", name)
}

// greet is declared with a parameter
//when calling greet, pass in as an argument
