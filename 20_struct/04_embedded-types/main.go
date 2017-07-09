package main

import (
	"fmt"
)

type Person struct {
	First string
	Last  string
	Age   int
}

type DoubleZero struct {
	Person
	LicenceToKill bool
}

func main() {
	p1 := DoubleZero{
		Person: Person{
			First: "James",
			Last:  "Bond",
			Age:   20},
		LicenceToKill: true,
	}
	p2 := DoubleZero{
		Person: Person{
			First: "Miss",
			Last:  "Moneypenny",
			Age:   18},
		LicenceToKill: false,
	}

	fmt.Println(p1.First, p1.Last, p1.Age, p1.LicenceToKill)
	fmt.Println(p2.First, p2.Last, p2.Age, p2.LicenceToKill)
}
