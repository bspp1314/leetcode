package main

import "fmt"

func totalNQueens(n int) int {
	if n == 1 {
		return 1
	}else if n == 2 {
		return 0
	}else if n == 3 {
		return 0
	}

	board := make([][]byte, n)
	for i := 0; i < n; i++ {
		board[i] =  make([]byte,n)
		for j := 0; j < len(board[i]); j++ {
			board[i][j]  = '.'
		}
	}

	res := 0
	backtrack(board,0,&res,n)

	return res
}

func backtrack(board [][]byte,row int,result *int ,n int)  {
	if row >=  n {
		*result = *result + 1
		return
	}

	// 列
	for col := 0; col < n ; col++ {
		if !check(board,row,col) {
			continue
		}

		board[row][col] = 'Q'
		backtrack(board,row+1,result,n)
		board[row][col] = '.'

	}
}

func check(queens [][]byte,row,col int) bool  {
	// 排除之前行的相同列
	for i := 0; i < row; i++ {
		if queens[i][col] == 'Q' {
			return false
		}
	}

	//排序之前行左上对角线
	for i, j := row-1, col-1; i >=0 && j >= 0; i, j = i-1, j-1 {
		if queens[i][j] == 'Q' {
			return false
		}
	}


	// 排除之前行右上对角线
	for i, j := row-1, col+1; i >=0 && j < len(queens); i, j = i-1, j+1 {
		if queens[i][j] == 'Q' {
			return false
		}
	}

	return true
}



func main() {
	res := totalNQueens(8)
	fmt.Println(res)
}
