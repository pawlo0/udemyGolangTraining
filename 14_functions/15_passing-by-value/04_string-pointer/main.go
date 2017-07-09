package main

import (
	"fmt"
)

func main() {
	name := "Todd"

	fmt.Println(name) // Todd

	changeMe(&name)

	fmt.Println(name) // Rocky
}

func changeMe(z *string) {
	fmt.Println(z)  // 0xc04200c290
	fmt.Println(*z) // Todd

	*z = "Rocky"

	fmt.Println(z) // 0xc04200c290
	fmt.Println(z) // Rocky
}
