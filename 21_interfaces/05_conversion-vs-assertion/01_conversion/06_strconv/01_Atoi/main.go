package main

import (
	"fmt"
	"strconv"
)

func main() {
	x := "12"
	y := 6
	z, _ := strconv.Atoi(x)
	fmt.Println(y + z)
}
