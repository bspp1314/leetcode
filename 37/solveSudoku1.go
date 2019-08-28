package main

import (
	"fmt"
)

func solveSudoku(board [][]byte) {
	solveSudokuHelp(board, 0, 0)

}
func solveSudokuHelp(board [][]byte, i, j int) bool {
	if i == 9 {
		return true
	}
	if j >= 9 {
		return solveSudokuHelp(board, i+1, 0)
	}
	if board[i][j] != '.' {
		return solveSudokuHelp(board, i, j+1)
	}

	var c byte
	for c = '1'; c <= '9'; c++ {
		if !isValid(board, i, j, c) {
			continue
		} else {
			board[i][j] = c
			if solveSudokuHelp(board, i, j+1) {
				return true
			}
			board[i][j] = '.'
		}
	}
	return false

}
func isValid(board [][]byte, i, j int, val byte) bool {
	for k := 0; k < 9; k++ {
		if board[k][j] == val {
			return false
		}
		if board[i][k] == val {
			return false
		}

		row := i/3*3 + k%3
		col := j/3*3 + k/3
		if board[row][col] == val {
			return false
		}
	}
	return true
}

func printSodu(in [][]byte) {
	for i := 0; i < 9; i++ {
		if i%3 == 0 {
			fmt.Println("=============================")
		}
		for j := 0; j < 9; j++ {
			if j%3 == 0 {
				fmt.Printf("||")
			}
			if in[i][j] == '.' {
				fmt.Printf(" . ")
			} else {
				fmt.Printf(" %d ", int(in[i][j]-'0'))
			}
		}
		fmt.Println()
	}
}
func main() {
	in := [][]byte{
		{'5', '3', '.', '.', '.', '.', '.', '.', '.'},
		{'6', '.', '.', '1', '9', '5', '.', '.', '.'},
		{'.', '9', '8', '.', '.', '.', '.', '6', '.'},
		{'8', '.', '.', '.', '6', '.', '.', '.', '3'},
		{'4', '.', '.', '8', '.', '3', '.', '.', '1'},
		{'7', '.', '.', '.', '2', '.', '.', '.', '6'},
		{'.', '6', '.', '.', '.', '.', '2', '8', '.'},
		{'.', '.', '.', '4', '1', '9', '.', '.', '5'},
		{'.', '.', '.', '.', '8', '.', '.', '7', '9'},
	}
	//printSodu(in)
	solveSudoku(in)
	printSodu(in)
	//fmt.Println(isValidSudoku(in))
}
