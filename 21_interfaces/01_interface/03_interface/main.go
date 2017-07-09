package main

import (
	"fmt"
	"math"
)

// First shape
type Square struct {
	side float64
}

// second shape
type Circle struct {
	radius float64
}

// function area connect to the Square struc
func (z Square) area() float64 {
	return z.side * z.side
}

// function area connect to the Circle struc
func (c Circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

// Then we implement the shape interface
type Shape interface {
	area() float64
}

// Now in info we can pass any shape,
// Since we can pass a Circle or a Squase
// When it call area for the Square object, it run area() function that is connected to the Squase struc
// When we call area for a Circle object, Go knows and it call the area() function that is connected to the Circle struc
func info(z Shape) {
	fmt.Println(z)
	fmt.Println(z.area())
}

func main() {
	s := Square{10}
	c := Circle{5}
	info(s)
	info(c)
}

// And we didn't changed much from the last example,
// Just added the Circle struct and the corresponding function
// But the info function is the same,
// because it runs the same interface
// which can take different objects that have the same signature, ie the same method.
