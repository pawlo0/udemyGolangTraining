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

	allPuzzles := getPuzzlesInArray()
	start := time.Now()
	count := 0
	for _, puzzle := range allPuzzles {
		fmt.Println(solve(puzzle))
		count++
		fmt.Println("Solved: ", count, " puzzles in:  ", time.Since(start))
	}
}

func solve(rawpuzzle []int) []int {
	puzzle, possibilities := initPuzzle(rawpuzzle)

	for cell := 0; cell < len(puzzle); cell++ {
		foundPossibility := false
		if len(possibilities[cell]) > 1 {
			for _, possibility := range possibilities[cell] {
				if possibility > puzzle[cell] && testThisNumber(possibility, cell, puzzle) {
					puzzle[cell] = possibility
					//fmt.Println("set: ", possibility, " in cell: ", cell)
					//fmt.Println(puzzle)
					//bufio.NewReader(os.Stdin).ReadBytes('\n')
					foundPossibility = true
					break
				}
			}
		} else {
			continue
		}
		if !foundPossibility {
			//fmt.Println("Trouble in cell: ", cell)
			//bufio.NewReader(os.Stdin).ReadBytes('\n')
			cell--
			for cell >= 0 {
				if len(possibilities[cell]) == 1 {
					cell--
				} else {
					if possibilities[cell][len(possibilities[cell])-1] > puzzle[cell] {
						stillPossibilities := false
						for _, possibility := range possibilities[cell] {
							if possibility > puzzle[cell] && testThisNumber(possibility, cell, puzzle) {
								stillPossibilities = true
								puzzle[cell] = possibility
								break
							}
						}
						if stillPossibilities {
							break
						} else {
							puzzle[cell] = 0
							cell--
						}
					} else {
						puzzle[cell] = 0
						//fmt.Println("Going to where should change cell to zero: ", puzzle[cell])
						//bufio.NewReader(os.Stdin).ReadBytes('\n')
						cell--
					}
				}
			}
			//fmt.Println("Going to try another number in cell: ", cell+1)
			//fmt.Println(puzzle)
			//bufio.NewReader(os.Stdin).ReadBytes('\n')
		}
	}

	return puzzle
}

// Initiates the puzzle by going throught all cells
// and find all possible alternatives for each cell.
// If there is only one alternative for one cell in this first round
// then that will be a numer already set for the final puzzle.
// also defines the cells that are initialy fixed numbers.
// So it returns 2 arrays, the puzzles itself, the array within array with alternatives
func initPuzzle(puzzle []int) ([]int, [][]int) {
	possibilities := [][]int{}
	for i, v := range puzzle {
		possibleForThisCell := []int{}
		if v == 0 {
			for number := 1; number < 10; number++ {
				if testThisNumber(number, i, puzzle) {
					possibleForThisCell = append(possibleForThisCell, number)
				}
			}
			if len(possibleForThisCell) == 1 {
				puzzle[i] = possibleForThisCell[0]
			}
		} else {
			possibleForThisCell = append(possibleForThisCell, v)
		}
		possibilities = append(possibilities, possibleForThisCell)
	}

	return puzzle, possibilities
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

func contains(s []int, e int) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

func getPuzzlesInArray() [][]int {

	url := "http://staffhome.ecm.uwa.edu.au/~00013890/sudoku17"

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
	/*

		var puzzle []int
		var allPuzzles [][]int

		lineStr := "000000010400000000020000000000050407008000300001090000300400200050100000000806000"

		tempLine := strings.Split(lineStr, "")
		for _, valueStr := range tempLine {
			value, _ := strconv.Atoi(valueStr)
			puzzle = append(puzzle, value)
		}
		allPuzzles = append(allPuzzles, puzzle)
		return allPuzzles
	*/
}
