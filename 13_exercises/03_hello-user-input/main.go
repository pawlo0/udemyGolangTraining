package main

import (
	"fmt"
)

func main() {
	var myName string
	fmt.Print("Enter your name: ")
	fmt.Scan(&myName)
	fmt.Println("Hello", myName)
}
