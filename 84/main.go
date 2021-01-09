package main

import (
	"fmt"
)



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

func largestRectangleArea2(heights []int) int {
	if len(heights) == 0 {
		return 0
	}

	n := len(heights)
	left, right := make([]int, n), make([]int, n)
	for i := 0; i < n; i++ {
		right[i] = n
	}

	stack := make([]int,0)


	for i := 0; i < len(heights); i++ {
		for len(stack) > 0 && heights[stack[len(stack)-1]] >= heights[i] {
			right[stack[len(stack)-1]]= i
			stack = stack[:len(stack)-1]
		}

		if len(stack) == 0 {
			left[i] = -1
		}else{
			left[i] = stack[len(stack)-1]
		}

		stack = append(stack,i)
	}

	max := 0
	for i := 0; i < n; i++ {
		max = Max(max,(right[i]-left[i]-1) * heights[i])
	}

	return max

}
func main() {
	fmt.Println(largestRectangleArea([]int{5,10,15,12}))
	fmt.Println(largestRectangleArea2([]int{5,10,15,12}))
}
