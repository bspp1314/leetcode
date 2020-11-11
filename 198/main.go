package main

import "fmt"

func rob(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	if len(nums) == 1 {
		return nums[0]
	}

	max := func(a,b int) int {
		if a > b {
			return a
		}

		return b
	}

	dp0 := nums[0]
	dp1 := max(nums[0],nums[1])

	for i := 2; i < len(nums); i++ {
		new := max(dp0+nums[i],dp1)
		dp0 = dp1
		dp1 = new

	}

	return dp1
}


func main() {
	out := rob([]int{1,1,3,6,7,10,7,1,8,5,9,1,4,4,3})
	fmt.Println(out)
}
