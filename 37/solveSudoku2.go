package main

import (
	"fmt"
)

func solveSudoku(board [][]byte) {
	row := make([][]bool, 9)
	col := make([][]bool, 9)
	cel := make([][]bool, 9)
	for i := 0; i < 9; i++ {
		row[i] = make([]bool, 10)
		col[i] = make([]bool, 10)
		cel[i] = make([]bool, 10)
	}
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if board[i][j] != '.' {
				num := board[i][j] - '0'
				row[i][num] = true
				col[j][num] = true
				cel[i/3*3+j/3][num] = true
			}
		}
	}
	solveSudokuHelp(board, 0, 0, row, col, cel)
}

func solveSudokuHelp(board [][]byte, i, j int, row, col, cel [][]bool) bool {
	for board[i][j] != '.' {
		if j++; j >= 9 {
			i++
			j = 0
		}
		if i >= 9 {
			return true
		}
	}
	var c byte
	for c = '1'; c <= '9'; c++ {

		celIndex := i/3*3 + j/3
		num := int(c - '0')
		if !row[i][num] && !col[j][num] && !cel[celIndex][num] {
			board[i][j] = c
			num := c - '0'
			row[i][num] = true
			col[j][num] = true
			cel[celIndex][num] = true
			if solveSudokuHelp(board, i, j, row, col, cel) {
				return true
			} else {
				row[i][num] = false
				col[j][num] = false
				cel[celIndex][num] = false
				board[i][j] = '.'
			}
		}

	}
	return false
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
