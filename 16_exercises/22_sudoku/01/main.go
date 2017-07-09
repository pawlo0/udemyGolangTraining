package main

import (
	"bufio"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"
)

func main() {
	allPuzzles := getPuzzlesInArray("https://projecteuler.net/project/resources/p096_sudoku.txt")
	start := time.Now()
	sum := 0
	for _, puzzle := range allPuzzles {
		solvedPuzzle := solve(puzzle)
		sum += solvedPuzzle[0]*100 + solvedPuzzle[1]*10 + solvedPuzzle[2]
	}
	fmt.Println(sum, " in ", time.Since(start))
}

func solve(puzzle []int) []int {

	// get index of fixed numbers
	var fixed []int
	for i, v := range puzzle {
		if v != 0 {
			fixed = append(fixed, i)
		}
	}
	// cycle the puzzle up
	for i := 0; i < len(puzzle); i++ {
		foundValue := false
		if !contains(fixed, i) {
			for value := puzzle[i] + 1; value < 10; value++ {
				if testThisNumber(value, i, puzzle) {
					puzzle[i] = value
					foundValue = true
					break
				}
			}
		} else {
			continue
		}

		if !foundValue {
			i--
			for i >= 0 {
				if contains(fixed, i) {
					i--
				} else if availableNumbers(puzzle[i]+1, i, puzzle) {
					i--
					break
				} else {
					puzzle[i] = 0
					i--
				}

			}
		}
	}
	return puzzle
}

func availableNumbers(start int, index int, puzzle []int) bool {
	for value := start; value < 10; value++ {
		if testThisNumber(value, index, puzzle) {
			return true
		}
	}
	return false
}

func contains(s []int, e int) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

func testThisNumber(n int, index int, puzzle []int) bool {

	for i, v := range puzzle {
		// test line
		if index/9 == i/9 && v == n {
			return false
		}
		// test column
		if index%9 == i%9 && v == n {
			return false
		}
		// test square
		if whichSquare(index) == whichSquare(i) && v == n {
			return false
		}
	}

	return true
}

func whichSquare(position int) int {
	square := 0
	triplet := position / 3
	if triplet > 8 && triplet < 18 {
		square += 3
	} else if triplet >= 18 {
		square += 6
	}
	square += triplet % 3
	return square
}

func getPuzzlesInArray(url string) [][]int {

	res, _ := http.Get(url)

	var rawLines []string

	subArray := bufio.NewScanner(res.Body)
	defer res.Body.Close()

	subArray.Split(bufio.ScanLines)

	for subArray.Scan() {
		rawLines = append(rawLines, subArray.Text())
	}

	var puzzle []int
	var allPuzzles [][]int
	count := 0
	for i, lineStr := range rawLines {
		if lineStr[0] != 71 {
			tempLine := strings.Split(lineStr, "")
			for _, valueStr := range tempLine {
				value, _ := strconv.Atoi(valueStr)
				puzzle = append(puzzle, value)
			}
			count++
			if count == 9 && i != 0 {
				allPuzzles = append(allPuzzles, puzzle)
				puzzle = []int{}
				count = 0
			}
		}
	}

	return allPuzzles
}
