package main

import "fmt"

func spiralOrder(matrix [][]int) []int {
	if len(matrix) == 0 {
		return []int{}
	}

	top := 0
	down := len(matrix) - 1
	left := 0
	right := len(matrix[0]) - 1

	out := make([]int, 0, len(matrix)*len(matrix[0]))
	spiralOrderHelp(matrix, left, right, top, down, &out)
	return out
}

func spiralOrder2(matrix [][]int) []int {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return []int{}
	}

	top := 0
	bottom := len(matrix) - 1
	left := 0
	right := len(matrix[0]) - 1

	res := make([]int,len(matrix)*len(matrix[0]))
	index := 0
	for left <= right && top <= bottom {
		for column := left; column <=right; column++ {
			res[index] = matrix[top][column]
			index++
		}

		for row := top+1;row <= bottom;row++ {
			res[index] = matrix[row][right]
			index++
		}

		if (left < right && top < bottom) {
			for column := right - 1; column > left; column-- {
				res[index] = matrix[bottom][column]
				index++
			}
			for row := bottom; row > top; row-- {
				res[index] = matrix[row][left]
				index++
			}
		}

		left++
		right--
		top++
		bottom--

	}

	return  res
}


func spiralOrderHelp(matrix [][]int, left int, right int, top int, down int, out *[]int) {
	if left > right {
		return
	}

	if top > down {
		return
	}

	if left == right {
		for i := top; i <= down; i++ {
			*out = append(*out, matrix[i][right])
		}
		return
	}

	if top == down {
		for i := left; i <= right; i++ {
			*out = append(*out, matrix[top][i])
		}
		return
	}

	for i := left; i < right; i++ {
		*out = append(*out, matrix[top][i])
	}

	for i := top; i < down; i++ {
		*out = append(*out, matrix[i][right])
	}

	for i := right; i > left; i-- {
		*out = append(*out, matrix[down][i])
	}

	for i := down; i > top; i-- {
		*out = append(*out, matrix[i][left])
	}

	spiralOrderHelp(matrix, left+1, right-1, top+1, down-1, out)
}

func main() {
	matrix := [][]int{
		{1, 2, 3, 10},
		{4, 5, 6, 11},
		{7, 8, 9, 12},
	}
	out := spiralOrder(matrix)
	for i := 0; i < len(out); i++ {
		fmt.Printf("%d ", out[i])
	}
	fmt.Println()
	fmt.Println("vim-go")

	spiralOrder2([][]int{{2,3}})
	spiralOrder2([][]int{{2},{3}})
}
