package main

import "fmt"

func main() {

	// you can only do this inside a func
	message := "Hello World!"
	a, b, c := 1, false, 3
	d := 4
	e := true

	fmt.Printf("%v \t %v \t %v \t %v \t %v \t %v \n", message, a, b, c, d, e)
	fmt.Printf("%T \t\t %T \t %T \t %T \t %T \t %T \n", message, a, b, c, d, e)
}
