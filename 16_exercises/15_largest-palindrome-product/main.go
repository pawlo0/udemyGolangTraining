package main

import (
	"fmt"
	"strconv"
)

func main() {
	maxValue := []int{0, 100, 100}
	for i := 100; i < 1000; i++ {
		for j := 100; j < 1000; j++ {
			if isPalindrome(i*j) && i*j > maxValue[0] {
				maxValue[0] = i * j
				maxValue[1] = i
				maxValue[2] = j
			}
		}
	}
	fmt.Println(maxValue[0], " ", maxValue[1], " ", maxValue[2])
}

func isPalindrome(value int) bool {
	val := strconv.Itoa(value)
	if len(val)%2 != 0 {
		return false
	}
	for i := 0; i <= len(val)/2; i++ {
		if val[i] == val[len(val)-(i+1)] {
			continue
		} else {
			return false
		}
	}
	return true
}
