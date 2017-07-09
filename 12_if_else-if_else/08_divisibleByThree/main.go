package main

import "fmt"

func main() {
	for i := 0; i <= 100; i++ {
		if statement := "is divisible by 3"; i%3 == 0 {
			fmt.Println(i, statement)
		}
	}
}
