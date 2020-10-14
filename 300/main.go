package main

import "fmt"

func main() {
	out := lengthOfLIS([]int{1, 2, 3, 4, 5})
	fmt.Println(out)
}

func lengthOfLIS(nums []int) int {
	if len(nums) <= 1 {
		return len(nums)
	}

	max := func(a, b int) int {
		if a >  b {
			return a
		}

		return b
	}

	// begin end
	dp := make([]int, len(nums))
	dp[0] = 1
	res := 1
	for i := 1; i < len(nums); i++ {
		dp[i] = 1
		for j := 0; j < i; j++ {
			if nums[i] > nums[j] {
				dp[i] = max(dp[i], dp[j]+1)
			}

		}

		res = max(res, dp[i])
	}

	return res
}
