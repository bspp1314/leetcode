package main

import "fmt"

func main() {
	out := lengthOfLIS2([]int{0,8,4,12,2})
	fmt.Println(out)
}


func lengthOfLIS2(nums []int) int {
	if len(nums) <= 1 {
		return len(nums)
	}

	smallestTails := []int{nums[0]}
	index := 0

	for i := 1; i < len(nums); i++ {
		if nums[i] > smallestTails[index] {
			smallestTails = append(smallestTails,nums[i])
			index++
		}else{
			left := 0
			right := len(smallestTails) -1
			mid := left + (right-left)/2

			for left < right {
				if nums[i] > smallestTails[mid] {
					left = mid+1
				} else if nums[i] < smallestTails[mid] {
					right = mid
				} else {
					right=mid
					break
				}
				mid = (left + right) / 2
			}

			smallestTails[right] = nums[i]
		}

		fmt.Println(smallestTails)
	}

	return len(smallestTails)
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

	// dp[i]为
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
