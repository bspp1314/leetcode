package main

import "fmt"

func generateMatrixHelp(matrix [][]int, left int, right int, top int, down int, key *int) {
	if left > right {
		return
	}

	if top > down {
		return
	}

	if left == right {
		for i := top; i <= down; i++ {
			matrix[i][right] =  *key
			*key = *key + 1
		}
		return
	}

	if top == down {
		for i := left; i <= right; i++ {
			matrix[top][i] = *key
			*key = *key + 1
		}
		return
	}

	for i := left; i < right; i++ {
		matrix[top][i] = *key
		*key = *key + 1
	}

	for i := top; i < down; i++ {
		matrix[i][right] = *key
		*key = *key + 1
	}

	for i := right; i > left; i-- {
		matrix[down][i] = *key
		*key = *key + 1
	}

	for i := down; i > top; i-- {
		matrix[i][left] = *key
		*key = *key + 1
	}

	generateMatrixHelp(matrix, left+1, right-1, top+1, down-1, key)
}

func generateMatrix(n int) [][]int {
	matrix := make([][]int,n)
	for i := 0; i < n; i++ {
		matrix[i] = make([]int,n)
	}

	key := 1

	top := 0
	down := len(matrix) - 1
	left := 0
	right := len(matrix[0]) - 1

	generateMatrixHelp(matrix,left,right,top,down,&key)

	return matrix
}

func main()  {
	fmt.Println(generateMatrix(5))
}
