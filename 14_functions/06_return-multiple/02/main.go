package main

import (
	"fmt"
)

func main() {
	s1, s2 := greet("Jane", "Doe")
	fmt.Println(s1)
	fmt.Println(s2)
}

func greet(fname, lname string) (string, string) {
	return fmt.Sprint(fname, lname), fmt.Sprint(lname, fname)
}
