package main

import "fmt"

func Max(a, b int) int {
	if a > b {
		return a
	}

	return b
}

func largestRectangleArea(heights []int) int {
	if len(heights) == 0 {
		return 0
	}

	stack := make([]int,0)

	res := 0
	heights = append(heights,0)

	for i := 0; i < len(heights); i++ {
		for len(stack) > 0 && heights[stack[len(stack)-1]] >= heights[i] {
			current := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			row := 0
			if len(stack) <= 0 {
				row = i
			} else {
				row = i - stack[len(stack)-1] - 1
			}

			res = Max(heights[current]*row, res)
		}

		stack = append(stack,i)
	}

	return res

}


func maximalRectangle(matrix [][]byte) int {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return 0
	}

	m := len(matrix)
	n := len(matrix[0])
	newRect := make([]int,n)
	max := 0

	for i := 0; i < m; i++ {
		for j :=0;j<n;j++ {
			if matrix[i][j] == '1' {
				newRect[j] += 1
			}else{
				newRect[j] = 0
			}


		}

		max = Max(largestRectangleArea(newRect),max)
	}

	return max
}

func main() {
	fmt.Println(maximalRectangle([][]byte{[]byte{'0','1'},{'1','0'}}))
}
