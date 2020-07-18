package main

import "fmt"

func rotate(matrix [][]int) {
	l := len(matrix)
	l2 := len(matrix) / 2
	for i := 0; i < l2; i++ {
		for j := 0; j < l; j++ {
			matrix[i][j], matrix[l-i-1][j] = matrix[l-i-1][j], matrix[i][j]
		}
	}

	for i := 0; i < l; i++ {
		for j := i; j < l; j++ {
			matrix[i][j], matrix[j][i] = matrix[j][i], matrix[i][j]
		}
	}

}
func main() {
	matrix := [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}
	rotate(matrix)
	for i := 0; i < len(matrix); i++ {
		fmt.Println(matrix[i])
	}
}
