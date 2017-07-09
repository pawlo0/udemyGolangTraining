// This although it works, it would take an eternity to find the solution for 49000 puzzles.
// In order to reduce the number of "testing" instead of full recursive backtesting,
// tried to first find the easy numbers, ie try to fill first the cells with one possibility.
// But found this would slow down even more the program.
// With that approach, what I was doing was going through every cell and "testing" possibilies
// which resolted in even more testing, that's why it was slower.

// My next approach would be to leave the recursive backtesting, ie start the same way as in here,
// but take the cycle where I do the testing and store the possible numbers in a map
// for the every empty cell. These possible numbers change with the recursive backtesting, but
// the result of the each test would return a limited number of possibilities and moving foward
// I would only then test the possible numbers instead of all numbers.
// I wouldn't even need test the possible numbers, if they are possible, I would just pick the first
// run another test, test will obviously pass, but would reduce the number of possibilities from test to test
// thus reducing the number of tests needed.

// Basicly would be a recursive backtesting but instead of pass or fails like I do now
// would be testing until only possibility left for that cell.
// ie, it would move forward with the first possible number, test next cell, move foward again with first possible number,
// and this number now may different than before the first test, but for sure it reduced the number of possibilities,
// and in each test return false if it finds any cell without a possible number.
// The diference is that this way the test returns not only if that number is good or not for that cell,
// but returns the possibilities and it fails if any cell has not possible number with the puzzle as it stands,
// forcing backward correction.
// move foward would narrow the number of possibilities, making process faster.
// move backwords would widen the number of possibilities, but it has to be in order to correct wrong numbers.

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
	start := time.Now()
	allPuzzles := getPuzzlesInArray("http://staffhome.ecm.uwa.edu.au/~00013890/sudoku17")
	fmt.Println("Time to prepare: ", time.Since(start))
	afterStart := time.Now()
	sum, n := 0, 0
	for _, puzzle := range allPuzzles {
		solvedPuzzle := solve(puzzle)
		sum += solvedPuzzle[0]*100 + solvedPuzzle[1]*10 + solvedPuzzle[2]
		n++
		fmt.Println("Solved ", n, " puzzles in ", time.Since(afterStart))
	}

	fmt.Println("Solved all in: ", time.Since(start))

}

func solve(puzzle []int) []int {

	// First easy solve the puzzle
	//puzzle = easySolve(puzzle)

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

// returns the same puzzle, but with positions with only one possibility already filed
func easySolve(puzzle []int) []int {
	for i, v := range puzzle {
		possibleNumbers := []int{}
		if v != 0 {
			for number := 1; number < 10; number++ {
				if testThisNumber(number, i, puzzle) {
					possibleNumbers = append(possibleNumbers, number)
				}
			}
			if len(possibleNumbers) == 1 {
				puzzle[i] = possibleNumbers[0]
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

	for _, lineStr := range rawLines {
		tempLine := strings.Split(lineStr, "")
		for _, valueStr := range tempLine {
			value, _ := strconv.Atoi(valueStr)
			puzzle = append(puzzle, value)
		}
		allPuzzles = append(allPuzzles, puzzle)
		puzzle = []int{}
	}

	return allPuzzles
}
