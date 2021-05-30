package main

import (
	"fmt"
	"math"
)

func jump(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	jumpMap := make(map[int]int)

	var dfs func(begin int, lastIndex int) int

	dfs = func(begin int, lastIndex int) int {
		if begin >= lastIndex {
			jumpMap[lastIndex] = 0
			return 0
		}

		if v, ok := jumpMap[begin]; ok {
			return v
		}

		min := math.MaxInt32
		for j := 1; j <= nums[begin]; j++ {
			v := dfs(begin+j, lastIndex)
			if v < min {
				min = v
			}
		}

		jumpMap[begin] = min + 1

		return min + 1
	}

	dfs(0, len(nums)-1)

	return jumpMap[0]
}

func main() {
	fmt.Println(jump([]int{2, 3, 1, 1, 4}))
}
