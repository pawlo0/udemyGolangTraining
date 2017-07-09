package main

import (
	"fmt"
)

func zero(x *int) {
	fmt.Println("address in zero", x)
	*x = 0
}

func main() {
	x := 5
	fmt.Println("address in main", &x)
	zero(&x)
	fmt.Println("End result:", x) // x is still 5
}
