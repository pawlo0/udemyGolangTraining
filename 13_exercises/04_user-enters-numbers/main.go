package main

import "fmt"

func main() {
	var large, small int
	fmt.Print("Enter a large number: ")
	fmt.Scan(&large)
	fmt.Print("Enter a number smaller than ", large, ":")
	fmt.Scan(&small)
	remainder := large % small
	fmt.Println("The remainder of", large, "divided by", small, "is", remainder)
}
