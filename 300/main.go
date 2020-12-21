package main

import "fmt"

func main() {
	out := lengthOfLIS([]int{0,4,8,2,4,5 })
	fmt.Println(out)
}

func lengthOfLIS(nums []int) int {
	if len(nums) <= 1 {
		return len(nums)
	}



	tails := []int{nums[0]}
	index := 0

	for i := 1; i < len(nums); i++ {
		if nums[i] > tails[index] {
			tails = append(tails,nums[i])
			index++
		}else{
			left := 0
			right := len(tails) -1
			mid := left + (right-left)/2

			for left < right {
				if nums[i] > tails[mid] {
					left = mid+1
				} else if nums[i] < tails[mid] {
					right = mid
				} else {
					right=mid
					break
				}
				mid = (left + right) / 2
			}

			tails[right] = nums[i]
		}

		fmt.Println(tails)
	}

	return len(tails)
}

func lengthOfLIS2(nums []int) int {
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
