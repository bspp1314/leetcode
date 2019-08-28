package main

import (
	"fmt"
)

type key struct {
	line int
	row  int
}

func isValidSudoku(board [][]byte) bool {
	dataMaps := make([](map[int]bool), 27)
	for i := 0; i < 9; i++ {
		dataMaps[i] = make(map[int]bool)
		dataMaps[i+9] = make(map[int]bool)
		dataMaps[i+18] = make(map[int]bool)
	}
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if board[i][j] == '.' {
				continue
			}
			key := int(board[i][j] - '1')
			if dataMaps[i][key] || dataMaps[j+9][key] || dataMaps[(i/3)*3+(j/3)+18][key] {
				return false
			}
			dataMaps[i][key] = true
			dataMaps[j+9][key] = true
			dataMaps[(i/3)*3+(j/3)+18][key] = true
		}
	}
	return true
}
func main() {
	in := [][]byte{
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
	fmt.Println(isValidSudoku(in))
}
