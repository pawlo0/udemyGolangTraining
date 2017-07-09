package main

import (
	"fmt"
	"math/big"
	"strconv"
)

func main() {
	var number big.Int
	number.Exp(big.NewInt(2), big.NewInt(1000), nil)
	str := number.String()
	sum := 0
	for i := 0; i < len(str); i++ {
		number, _ := strconv.Atoi(string(str[i]))
		sum += number
	}
	fmt.Println(sum)
}
