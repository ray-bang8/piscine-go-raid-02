package main

import (
	"fmt"
	"os"
)

func main() {
	matrix := os.Args[1:]
	if !checkForValidity(matrix) {
		return
	}
	var m [9][9]int
	m = toArray(matrix)

	if !isSudokuSolved(m) {
		fmt.Println("Error")
	}
}

func isSudokuSolved(m [9][9]int) bool {

	var row, col int
	isFilled := true

	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if m[i][j] == 0 {
				row = i
				col = j
				isFilled = false
				break
			}
		}
		if !isFilled {
			break
		}
	}

	if isFilled {
		for i := 0; i < 9; i++ {
			for index, j := range m[i] {
				fmt.Print(j)
				if index != 8 {
					fmt.Print(" ")
				}
			}
			fmt.Println()
		}

		return true
	}

	for n := 1; n <= 9; n++ {
		if isAllowed(m, row, col, n) {
			m[row][col] = n
			if isSudokuSolved(m) {
				return true
			}
			m[row][col] = 0
		}
	}
	return false
}

func isAllowed(m [9][9]int, row int, col int, num int) bool {
	for i := 0; i < 9; i++ {
		if m[row][i] == num || m[i][col] == num {
			return false
		}
	}

	sqrt := 3
	rowstart := row - row%sqrt
	colstart := col - col%sqrt

	for i := rowstart; i < rowstart+sqrt; i++ {
		for j := colstart; j < colstart+sqrt; j++ {
			if m[i][j] == num {
				return false
			}
		}
	}
	return true
}

func toArray(m []string) [9][9]int {
	var arr [9][9]int
	var count int
	for i, line := range m {
		for j, letter := range line {
			if letter == '.' {
				arr[i][j] = 0
			} else {
				count = 0
				for nb := '0'; nb < letter; nb++ {
					count++
				}
				arr[i][j] = count
			}
		}
	}
	return arr
}

func checkForValidity(s []string) bool {
	if len(s) != 9 {
		fmt.Println("Error")
		return false
	}
	for _, line := range s {
		if len(line) != 9 {
			fmt.Println("Error")
			return false
		}
		for _, letter := range line {
			if letter != '.' && (letter < '1' || letter > '9') {
				fmt.Println("Error")
				return false
			}
		}
	}
	return true
}
