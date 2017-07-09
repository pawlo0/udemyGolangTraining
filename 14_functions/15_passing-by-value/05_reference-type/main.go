package main

import (
	"fmt"
)

func main() {
	m := make(map[string]int)
	fmt.Println(m["Todd"]) // 0
	changeMe(m)
	fmt.Println(m["Todd"]) // 44
}

func changeMe(z map[string]int) {
	z["Todd"] = 44
}
