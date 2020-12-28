package main

import (
	"fmt"
	"log"
)

type NumMatrix struct {
	Dp [][]int //每一行的前缀和

}

func Constructor(matrix [][]int) NumMatrix {
	s := NumMatrix{Dp: nil}
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return s 
	}

	row := len(matrix)
	col := len(matrix[0])
	s.Dp = make([][]int,row+1)

	for i := 0; i < len(s.Dp); i++ {
			s.Dp[i] = make([]int,col+1)
	}
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			s.Dp[i + 1][j + 1] = s.Dp[i + 1][j] + s.Dp[i][j + 1] + matrix[i][j] - s.Dp[i][j]
		}
	}
	
	return s 
}


func (this *NumMatrix) SumRegion(row1 int, col1 int, row2 int, col2 int) int {
	if this.Dp == nil {
		return 0
	}

	return this.Dp[row2 + 1][col2 + 1] - this.Dp[row1][col2 + 1] - this.Dp[row2 + 1][col1] + this.Dp[row1][col1]

}

func main() {

	s := Constructor([][]int{
		{3, 0, 1, 4, 2},
		{5, 6, 3, 2, 1},
		{1, 2, 0, 1, 5},
		{4, 1, 0, 1, 7},
		{1, 0, 3, 0, 5},
	})

	log.Println(s.Dp)

	fmt.Println(s.SumRegion(2,1,4,3))
}