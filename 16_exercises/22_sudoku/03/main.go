package main

import (
	"fmt"
	"strconv"
	"strings"
)

type cell struct {
	value         int
	possibilities []int
	fixed         bool
}
type structPuzzle struct {
	cells     []cell
	puzzleArr []int
}

func main() {

	allPuzzles := getPuzzlesInArray()

	fmt.Println(solve(initPuzzle(allPuzzles[0]), 0))
}

// Initiates the puzzle struct by going throught all cells
// and find all possible alternatives.
// If there is only one alternative for one cell in this first round
// then that will be a numer already set for the final puzzle.
// also defines the cells thar are initialy fixed numbers.
func initPuzzle(puzzle []int) structPuzzle {
	cells := []cell{}
	for i, v := range puzzle {
		possibleForThisCell := []int{}
		fixed := false
		if v == 0 {
			for number := 1; number < 10; number++ {
				if testThisNumber(number, i, puzzle) {
					possibleForThisCell = append(possibleForThisCell, number)
				}
			}
			if len(possibleForThisCell) == 1 {
				v = possibleForThisCell[0]
				puzzle[i] = v
			}
		} else {
			fixed = true
			possibleForThisCell = append(possibleForThisCell, v)
		}
		newCell := cell{v, possibleForThisCell, fixed}
		cells = append(cells, newCell)
	}
	newStructPuzzle := structPuzzle{cells, puzzle}
	return newStructPuzzle
}

// Recursive function to solve the puzzle
// picks the initialized struct and goes through the cells
// advantage is that for each cell it doesn't run all the numbers
// but just the ones that are possible for that cell from the first initialization
// 1st step, picks the first alternative possible
// moves foward, until if doesn't find a cell where all possible alternatives doesn't work
// has to move back and try another possibility
// if from start to the cell where it is all cells have only one possibility,
// then the puzzle is closed until that cell.
func solve(puzzle structPuzzle, cell int) structPuzzle {

	if cell == 0 {
		return puzzle
	}
	if puzzle.cells[cell].fixed {
		cell++
		solve(puzzle, cell)
	}
	for _, v := range puzzle.cells[cell].possibilities {
		if v > puzzle.cells[cell].value && testThisNumber(v, cell, puzzle.puzzleArr) {
			puzzle.cells[cell].value = v
		}
		// se para trás todos os possibilities têm tamanho zero, então passa a fixo

	}
	return puzzle

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
	/*
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
	*/

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

}
