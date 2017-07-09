package main

import "fmt"

func main() {
	myGreeting := make(map[string]string)
	myGreeting["Todd"] = "Good morning!"
	myGreeting["Jenny"] = "Bonjour!"

	fmt.Println(myGreeting)
}
