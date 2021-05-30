package main

import "fmt"

func solveNQueens2(n int) [][]string {
	if n == 1 {
		return [][]string{[]string{"Q"}}
	}else if n == 2 {
		return [][]string{}
	}else if n == 3 {
		return [][]string{}
	}

	board := make([][]byte, n)
	for i := 0; i < n; i++ {
		board[i] =  make([]byte,n)
		for j := 0; j < len(board[i]); j++ {
			board[i][j]  = '.'
		}
	}

	var dfs func(row int)
	res := make([][]string,0)
	columns := map[int]bool{}
	diagonals1, diagonals2 := map[int]bool{}, map[int]bool{}




	dfs = func(row int) {
		if row >= n {
			res = append(res,generateBoard(board,n))
			return
		}
		// 列
		for col := 0; col < n ; col++ {
			if columns[col] {
				continue
			}

			diagonal1 := row - col
			if diagonals1[diagonal1] {
				continue
			}
			diagonal2 := row + col
			if diagonals2[diagonal2] {
				continue
			}

			columns[col] = true
			diagonals1[diagonal1] = true
			diagonals2[diagonal2] = true
			board[row][col] = 'Q'
			dfs(row+1)
			board[row][col] = '.'
			delete(columns,col)
			delete(diagonals1,diagonal1)
			delete(diagonals2,diagonal2)
		}
	}

	dfs(0)


	return  res


}

func solveNQueens(n int) [][]string {
	if n == 1 {
		return [][]string{[]string{"Q"}}
	}else if n == 2 {
		return [][]string{}
	}else if n == 3 {
		return [][]string{}
	}

	board := make([][]byte, n)
	for i := 0; i < n; i++ {
		board[i] =  make([]byte,n)
		for j := 0; j < len(board[i]); j++ {
			board[i][j]  = '.'
		}
	}

	res := make([][]string,0)
	backtrack(board,0,&res,n)

	return res
}

func backtrack(board [][]byte,row int,result *[][]string,n int)  {
	if row >=  n {
		*result = append(*result,generateBoard(board,n))
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

func generateBoard(queens [][]byte, n int) []string {
	var board []string
	for i := 0; i < n; i++ {
		tem := make([]byte,len(queens[i]))
		copy(tem,[]byte(queens[i]))
		board = append(board, string(tem))
	}
	return board
}


func main() {
	res := solveNQueens(3)

	for i := 0; i < len(res); i++ {
		for j := 0; j < len(res[i]); j++ {
			fmt.Println(res[i][j])
		}

		fmt.Println()
	}
}
