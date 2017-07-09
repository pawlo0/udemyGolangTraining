package main

import (
	"fmt"
)

func main() {
	switch "Medhi" {
	case "Daniel":
		fmt.Println("Messup Daniel")
	case "Medhi":
		fmt.Println("Messup Medhi")
	case "Jenny":
		fmt.Println("Messup Jenny")
	default:
		fmt.Println("Have you no friends")
	}
}
