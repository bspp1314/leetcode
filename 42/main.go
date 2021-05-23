package main

import (
	"fmt"
)

func min(a,b int) int  {
	if a < b {
		return a
	}

	return b
}

func trap2(height []int) int   {
	if len(height) <= 2 {
		return 0
	}
	var stack []int
	var ans int

	for i,h := range height {
		for len(stack) > 0 && h > height[stack[len(stack)-1]] {
			top := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			if len(stack) == 0 {
				break
			}

			left := stack[len(stack)-1]

			curWidth := i - left - 1
			curHeight := min(height[left], h) - height[top]
			ans += curWidth * curHeight

		}

		stack = append(stack, i)
	}



	return ans
}

func max(a,b int) int   {
	if a > b {
		return a
	}

	return b	
}

func trap(height []int) int {
	if len(height) <= 2 {
		return 0
	}
	leftMax := make([]int,len(height))
	leftMax[0] = 0
	for i := 1; i < len(height); i++ {
		leftMax[i] = max(leftMax[i-1],height[i-1])
	}

	rightMax := height[len(height) -1 ]
	res := 0
	for i:= len(height) -2;i >= 1;i-- {
		if height[i] >=  rightMax || height[i] >= leftMax[i] {
			rightMax = max(rightMax,height[i])
			continue
		}



		if leftMax[i] < rightMax {
			res +=leftMax[i]  - height[i]
		}else{
			res += rightMax -height[i]
		}

		rightMax = max(rightMax,height[i])

	}

	return res

}


func trap3(height []int) int {
	if len(height) <= 2 {
		return 0
	}
	leftMax := 0
	rightMax := 0
	res := 0
	left := 0
	right := len(height) - 1

	for left <= right {
		leftMax = max(leftMax,height[left])
		rightMax = max(rightMax,height[right])

		if leftMax < rightMax {
			res += leftMax - height[left]
			left++
		}else{
			res += rightMax - height[right]
			right--
		}
	}

	return res

}


func main() {
	fmt.Println(trap([]int{0,1,0,2,1,0,1,3,2,1,2,1}))
	fmt.Println(trap([]int{4,2,0,3,2,5}))
	fmt.Println(trap2([]int{0,1,0,2,1,0,1,3,2,1,2,1}))
	fmt.Println(trap2([]int{4,2,0,3,2,5}))
	fmt.Println(trap3([]int{0,1,0,2,1,0,1,3,2,1,2,1}))
	fmt.Println(trap3([]int{4,2,0,3,2,5}))
}
