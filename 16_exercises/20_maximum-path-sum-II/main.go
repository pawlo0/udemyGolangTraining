package main

import (
	"bufio"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"
)

// Confess that I was only able to find the answer after reading solutions of problem 18, maximum path sum part I
// Altough was able to find solution of problem 18 without brut force
// that didn't work for this larger triangle
// found this simple and beautiful solution in the forum of problem 18 and tried in this problem and it works wonderfully
// and super fast

func main() {
	start := time.Now()
	array := getArray("https://projecteuler.net/project/resources/p067_triangle.txt")
	fmt.Println(time.Since(start))
	start = time.Now()
	for row := len(array) - 2; row >= 0; row-- {
		for i := range array[row] {
			array[row][i] += max(array[row+1][i], array[row+1][i+1])
		}
	}
	fmt.Println(array[0][0], time.Since(start))

}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func getArray(url string) [][]int {

	res, _ := http.Get(url)

	var lines []string

	subArray := bufio.NewScanner(res.Body)
	defer res.Body.Close()

	subArray.Split(bufio.ScanLines)

	for subArray.Scan() {
		lines = append(lines, subArray.Text())
	}

	firstValue, _ := strconv.Atoi(lines[0])
	firstArray := []int{firstValue}
	array := [][]int{firstArray}
	for i := 1; i < len(lines); i++ {
		lineArrayOfStrings := strings.Split(lines[i], " ")
		var lineArrayOfInts []int
		for _, v := range lineArrayOfStrings {
			value, _ := strconv.Atoi(v)
			lineArrayOfInts = append(lineArrayOfInts, value)
		}
		array = append(array, lineArrayOfInts)
	}
	return array

}
