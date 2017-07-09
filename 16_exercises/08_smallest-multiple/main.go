package main

import (
	"fmt"
)

func main() {
	i := 11
	for {
		check := true
		for j := 1; j <= 20; j++ {
			if i%j != 0 {
				check = false
				break
			}
		}
		if check == true {
			fmt.Println(i)
			break
		}
		i++
	}
}
