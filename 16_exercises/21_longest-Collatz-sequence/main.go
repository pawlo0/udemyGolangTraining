package main

import (
	"fmt"
)

func main() {
	maxTerms, terms := 0, 0
	for i := 1000000; i > 100; i-- {
		terms = howMany(i)
		if terms > maxTerms {
			maxTerms = terms
			fmt.Println(i, maxTerms)
		}
	}

}

func howMany(start int) int {
	num := start
	terms := 1
	for num != 1 {
		if num%2 == 0 {
			num = num / 2
		} else {
			num = num*3 + 1
		}
		terms++
	}
	return terms
}
