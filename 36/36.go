package main

import "fmt"

func isValidSudoku(board [][]byte) bool {
	if len(board) <= 0 || len(board) != 9 || len(board[0]) != 9 {
		return false
	}
	row := [9][9]bool{{}}
	col := [9][9]bool{{}}
	cell := [9][9]bool{{}}
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if board[i][j] >= '1' && board[i][j] <= '9' {
				v := board[i][j] - '1'
				if row[i][v] || col[v][j] || cell[3*(i/3)+j/3][v] {
					return false
				}
				row[i][v] = true
				col[v][j] = true
				cell[3*(i/3)+j/3][v] = true
			}
		}
	}

	return true
}
func main() {
	m := [][]byte{
		{'5', '3', '.', '.', '7', '.', '.', '.', '.'},
		{'6', '.', '.', '1', '9', '5', '.', '.', '.'},
		{'.', '9', '8', '.', '.', '.', '.', '6', '.'},
		{'8', '.', '.', '.', '6', '.', '.', '.', '3'},
		{'4', '.', '.', '8', '.', '3', '.', '.', '1'},
		{'7', '.', '.', '.', '2', '.', '.', '.', '6'},
		{'.', '6', '.', '.', '.', '.', '2', '8', '.'},
		{'.', '.', '.', '4', '1', '9', '.', '.', '5'},
		{'.', '.', '.', '.', '8', '.', '.', '7', '9'},
	}
	fmt.Println(isValidSudoku(m))
	fmt.Println("vim-go")
}
