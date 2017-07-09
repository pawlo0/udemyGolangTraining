package main

import "fmt"

type foo int

func main() {
	var myAge foo
	myAge = 44
	fmt.Printf("%T %v \n", myAge, myAge)

	var yourAge int
	yourAge = 29
	fmt.Printf("%T %v \n", yourAge, yourAge)

	// This won't wort, can add two different types, although both are int
	//fmt.Println(myAge + yourAge)//

	// this works, because this conversion works
	fmt.Println(int(myAge) + yourAge)
}
