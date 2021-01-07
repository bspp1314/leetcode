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

	stack := make([]int, len(heights)+1)
	pointer := -1

	res := 0
	heights = append(heights,0)

	for i := 0; i < len(heights); i++ {
		for pointer >= 0 && heights[stack[pointer]] >= heights[i] {
			current := stack[pointer]
			pointer--
			row := 0
			if pointer < 0 {
				row = i
			} else {
				row = i - stack[pointer] - 1
			}


			res = Max(heights[current]*row, res)
		}

		pointer++
		stack[pointer] = i
		fmt.Println(stack)
	}

	return res

}
func main() {
	fmt.Println(largestRectangleArea([]int{2, 4,8,5}))
	//fmt.Println(largestRectangleArea([]int{2,1,5,6,2,3}))
}
