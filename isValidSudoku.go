package main

import (
	"fmt"
)

type key struct {
	line int
	row  int
}

func isValidSudoku(board [][]byte) bool {
	dataMaps := make([](map[byte]bool), 27)
	for i := 0; i < 9; i++ {
		dataMaps[i] = make(map[byte]bool)
		dataMaps[i+9] = make(map[byte]bool)
		dataMaps[i+18] = make(map[byte]bool)
	}
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if board[i][j] == '.' {
				continue
			}
			key := board[i][j] - '1'
			if dataMaps[i][key] || dataMaps[j+9][key] || dataMaps[(i/3)*3+(j/3)+18][key] {
				fmt.Println("i", i+1)
				fmt.Println("j", j+1)
				return false
			}
			dataMaps[i][key] = true
			dataMaps[j+9][key] = true
			dataMaps[(i/3)*3+(j/3)+18][key] = true
		}
	}
	return true
}

func isValidSudokuSim(board [][]byte, i, j int, val byte) bool {
	for k := 0; k < 9; k++ {
		//所在的行
		if board[i][k] == val {
			return false
		}
	}
	for k := 0; k < 9; k++ {
		//检查所在的列
		if board[k][j] == val {
			return false
		}
	}

	//检查所在的九宫格
	//	for k := 0; k < 9; k++ {
	//		row := i/3*3 + k/3
	//		col := j/3*3 + k%3
	//		if board[row][col] == val {
	//			return false
	//		}
	//	}
	//}
	//for k := 0; k < 9; k++ {
	//	if board[i][k] == val {
	//		if i == 0 && j == 2 {
	//			fmt.Println("val", int(val))
	//		}
	//		return false
	//	}

	//	if board[k][j] == val {
	//		if i == 0 && j == 2 {
	//			fmt.Println("val", int(val-'0'))
	//		}
	//		return false
	//	}
	//	row := i/3*3 + k/3
	//	col := j/3*3 + k%3
	//	if board[row][col] == val {
	//		if i == 0 && j == 2 {
	//			fmt.Println("row", row)
	//			fmt.Println("col", col)
	//			fmt.Println("val", int(val-'0'))
	//		}
	//		return false
	//	}
	//}
	//return true
	return true
}
func solveSudoku(board [][]byte) {
	illegalMaps := make([][]map[byte]bool, 9)
	for i := 0; i < 9; i++ {
		illegalMaps[i] = make([]map[byte]bool, 9)
		for j := 0; j < 9; j++ {
			illegalMaps[i][j] = make(map[byte]bool)
		}
	}
	solveSudokuHelp(board, 0, 0, illegalMaps)

}
func solveSudokuHelp(board [][]byte, i, j int, illegalMaps [][]map[byte]bool) {
	for ; i < 9; i++ {
		for ; j < 9; j++ {
			if board[i][j] != '.' {
				continue
			}

			var val byte
			needBack := true
			for val = '1'; val <= '9'; val++ {
				if illegalMaps[i][j][val] {
					continue
				} else {
					if isValidSudokuSim(board, i, j, val) {
						board[i][j] = val
						needBack = false
					}
				}
			}
			if needBack {
				if i == 0 && j == 0 {
					return
				} else if j == 0 {
					illegalMaps[i-1][9][board[i-1][9]] = true
					board[i-1][9] = '.'
					solveSudokuHelp(board, i-1, 9, illegalMaps)
				} else {
					illegalMaps[i][j-1][board[i][j-1]] = true
					board[i][j-1] = '.'
					solveSudokuHelp(board, i, j-1, illegalMaps)
				}
				fmt.Println(int(board[i][j] - '0'))
			}
		}
	}
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
		{'8', '3', '.', '.', '7', '.', '.', '.', '.'},
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
