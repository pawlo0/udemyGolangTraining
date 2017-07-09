package main

import "fmt"

func main() {
	myGreeting := map[string]string{
		"Todd":  "Good morning!",
		"Jenny": "Bonjour!",
	}

	fmt.Println(myGreeting["Jenny"])
}
