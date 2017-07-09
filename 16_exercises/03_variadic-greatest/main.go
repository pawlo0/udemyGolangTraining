package main

import "fmt"

func main() {
	fmt.Println(max(2, 5, 8, 444, 100, 333, 222))
}

func max(n ...int) int {
	var largest int
	for _, v := range n {
		if v > largest {
			largest = v
		}
	}
	return largest
}
