package main

import (
	"fmt"
	"math"
)

func findSubsequences(nums []int) [][]int {
	if len(nums) == 0 {
		return [][]int{}
	}


	var ans  [][]int
	dfs(nums,0,math.MinInt32,[]int{},&ans)
	return ans
}


func dfs(nums []int, left int, last int,temp []int,ans *[][]int) {
	if left == len(nums) {
		if len(temp) >= 2 {
			t := make([]int, len(temp))
			copy(t, temp)
			*ans = append(*ans, t)
		}

		return
	}


	if nums[left] >= last {
		k := append(temp, nums[left])
		dfs(nums, left+1,nums[left],k,ans)
	}
	if nums[left] != last {
		dfs(nums, left+1, last,temp,ans)
	}
}

func main() {
	fmt.Println(findSubsequences([]int{1,3,5,7}))
}
