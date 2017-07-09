package main

import (
	"fmt"
)

func main() {
	mySlice := []string{"a", "b", "c", "g", "m", "z"}
	fmt.Println(mySlice)
	fmt.Println(mySlice[2:4])  // slicing a slyce
	fmt.Println(mySlice[2])    // index access; accessing by index 2
	fmt.Println("myString"[2]) // index access; acessing by index
}
