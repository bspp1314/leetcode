package main

import "fmt"

// 1 ~ N+1

func firstMissingPositive(nums []int) int {
	for i := 0; i < len(nums); i++ {
		for nums[i] >= 1 && nums[i] <= len(nums) && nums[i] != nums[nums[i]-1]{
			nums[i],nums[nums[i]-1] = nums[nums[i]-1],nums[i]
		}
	}


	for i := 0;i<len(nums);i++ {
		if nums[i] != i+1 {
			return i+1
		}
	}

	return len(nums) +1
}

func main() {
	fmt.Println(firstMissingPositive([]int{-1,1,2,3,3,5,1}))
}
