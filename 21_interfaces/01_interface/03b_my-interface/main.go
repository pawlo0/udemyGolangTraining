package main

import (
	"fmt"
	"math"
)

type Square struct {
	side float64
}

func (z Square) area() float64 {
	return z.side * z.side
}

type Circle struct {
	radius float64
}

func (c Circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

type Rectangle struct {
	side1 float64
	side2 float64
}

func (r Rectangle) area() float64 {
	return r.side1 * r.side2
}

type Shape interface {
	area() float64
}

func info(z Shape) {
	fmt.Println(z)
	fmt.Println(z.area())
}

func main() {
	s := Square{10}
	c := Circle{5}
	r := Rectangle{5, 4}
	info(s)
	info(c)
	info(r)
}
