package main

import (
	"fmt"
	"math"
	"time"
)

func main() {
	start := time.Now()
	var a, b, c float64
	var finish bool
	for b = 1; b < 1000 && !finish; b++ {
		for a = 1; a < 1000 && a < b; a++ {
			c = math.Sqrt(a*a + b*b)
			if a+b+c == 1000 {
				fmt.Println(a, " ", b, " ", c, " ", a+b+c, " ", int(a*b*c))
				finish = true
				break
			}
		}
	}
	fmt.Println(time.Since(start))
}
