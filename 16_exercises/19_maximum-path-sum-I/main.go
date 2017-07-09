package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	triangle := "75\n95 64\n17 47 82\n18 35 87 10\n20 04 82 47 65\n19 01 23 75 03 34\n88 02 77 73 07 63 67\n99 65 04 28 06 16 70 92\n41 41 26 56 83 40 80 70 33\n41 48 72 33 47 32 37 16 94 29\n53 71 44 65 25 43 91 52 97 51 14\n70 11 33 28 77 73 17 78 39 68 17 57\n91 71 52 38 17 14 91 43 58 50 27 29 48\n63 66 04 68 89 53 67 30 73 16 69 87 40 31\n04 62 98 27 23 09 70 98 73 93 38 53 60 04 23"

	array := getArray(triangle)

	sum := array[0][0]

	path := []int{array[0][0]}

	var j int

	for i := 1; i < len(array); i++ {

		if i == len(array)-1 && array[i][j] < array[i][j+1] || i < len(array)-1 && array[i][j]+(array[i+1][j]+array[i+1][j+1])/2 < array[i][j+1]+(array[i+1][j+1]+array[i+1][j+2])/2 {
			j++
		}
		sum += array[i][j]
		path = append(path, array[i][j])
	}

	fmt.Println(path)
	fmt.Println(sum)
}

func getArray(s string) [][]int {
	subArray := strings.Split(s, "\n")
	firstValue, _ := strconv.Atoi(subArray[0])
	firstArray := []int{firstValue}
	array := [][]int{firstArray}
	for i := 1; i < len(subArray); i++ {
		lineArrayOfStrings := strings.Split(subArray[i], " ")
		var lineArrayOfInts []int
		for _, v := range lineArrayOfStrings {
			value, _ := strconv.Atoi(v)
			lineArrayOfInts = append(lineArrayOfInts, value)
		}
		array = append(array, lineArrayOfInts)
	}
	return array
}
