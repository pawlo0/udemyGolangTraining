package main

import (
	"fmt"
)

type Person struct {
	Name string
	Age  int
}

type DoubleZero struct {
	Person
	LicenceToKill bool
}

func (p Person) Greeting() {
	fmt.Println("I'm just a regular person.")
}

func (dz DoubleZero) Greeting() {
	fmt.Println("Miss Moneypenny, so good to see you.")
}

func main() {
	p1 := Person{"Ian Flemming", 44}
	p2 := DoubleZero{
		Person: Person{
			Name: "James",
			Age:  20},
		LicenceToKill: true,
	}

	p1.Greeting()
	p2.Greeting()
	p2.Person.Greeting()
}
