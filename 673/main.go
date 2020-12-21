package main

import (
	"fmt"
	"log"
)

func findNumberOfLIS(nums []int) int {
	if len(nums) <= 1 {
		return len(nums)
	}

	// dp[i]为
	dp := make([]int, len(nums))
	dpCount := make([]int,len(nums))
	for i := 0; i < len(dpCount); i++ {
		dp[i] = 1
		dpCount[i] = 1

	}

	max  := dp[0]

	for i := 1; i < len(nums); i++ {
		for j := 0; j < i ; j++ {
			if nums[i] > nums[j] {
				if dp[j]+1 > dp[i] {
					dp[i] = dp[j]+1
					dpCount[i] = dpCount[j]//更新为count[j]
				} else if dp[j]+1 == dp[i]{
					dpCount[i] += dpCount[j]//与count[j]相加
				}
			}
		}

		if max < dp[i] {
			max = dp[i]
		}
	}

	res := 0

	for i := 0; i < len(dp); i++ {
		if dp[i] == max {
			res+= dpCount[i]
		}
	}

	log.Println(dp)


	return res
}


func findNumberOfLIS2(nums []int) int {
	if len(nums) <= 1 {
		return len(nums)
	}

	tails := []int{nums[0]}
	index := 0
	endMax := make([]int, len(nums))
	dpCount := make([]int,len(nums))
	for i := 0; i < len(dpCount); i++ {
		endMax[i] = 1
		dpCount[i] = 1

	}

	for i := 1; i < len(nums); i++ {
		if nums[i] > tails[index] {
			tails = append(tails,nums[i])
			index++
			dpCount[i] = dpCount[i-1]
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

		endMax[i] = len(tails)

	}

	log.Println(endMax)

	return len(tails)
}

func main() {
	fmt.Println(findNumberOfLIS2([]int{1,3,5,4,7,9,8,9}))
	fmt.Println(findNumberOfLIS([]int{1,3,5,4,7,9,8,9}))
}
