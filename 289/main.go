package main

import "fmt"

func gameOfLife(board [][]int)  {
	m := len(board)
	if m == 0 {
		return
	}
	n := len(board[0])
	if n == 0 {
		return
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			help(i,j,m,n,board)
		}
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if board[i][j] == -1 {
				board[i][j] = 0
			}else if board[i][j] == 2 {
				board[i][j] = 1
			}
		}
	}

}

func help(i,j,m,n int,board [][]int)  {
	liveCount := 0
	deadCount := 0
	for row := i-1; row <=i+1 ; row++ {
		for col := j-1;col <= j+1; col++ {
			if row < 0 || row >= m || col < 0 || col >=n || (row == i && col == j ){
				continue
			}

			if board[row][col] == 1 || board[row][col] == -1 {
				liveCount++
			}else{
				deadCount++
			}
		}
	}

	if board[i][j] == 1 {
		if liveCount <= 1 || liveCount > 3 {
			board[i][j] = -1
		}
	}else{
		if liveCount ==3 {
			board[i][j] = 2
		}
	}
}

func PrintNum(num [][]int)  {
	fmt.Println("begin")
	for i := 0; i < len(num); i++ {
		fmt.Println(num[i])
	}
	fmt.Println("end")
}

func main() {
	num := [][]int{
		{0,1,0},
		{0,0,1},
		{1,1,1},
		{0,0,0},
	}
	gameOfLife(num)
	PrintNum(num)
}

