package main

import "fmt"

func main() {
	fmt.Println(lengthOfLIS([]int{0,4,8,2,4,5 }))
	//fmt.Println(lengthOfLIS([]int{0}))
	fmt.Println(lengthOfLIS([]int{10,9,2,5,3,4}))
	fmt.Println(lengthOfLIS([]int{4,10,4,3,8,9}))
	//fmt.Println(findFirstGe([]int{2,5},3))
	//fmt.Println(kk([]int{2,5},3))
}

func lengthOfLIS(nums []int) int {
	if len(nums) <= 1 {
		return len(nums)
	}

	tails := []int{nums[0]}
	for i := 1; i < len(nums); i++ {
		if nums[i] > tails[len(tails)-1] {
			tails = append(tails,nums[i])
		}else{
			left := 0
			right := len(tails) -1

			//第一个大于等于
			for left <= right {
				mid := left + (right-left) >> 1
				if tails[mid] < nums[i] {
					left = mid + 1
				}else if tails[mid] >= nums[i] {
					right = mid - 1
				}
			}
			tails[right+1] = nums[i]
		}
	}

	return len(tails)
}

func findFirstGe(nums []int,target int) int  {
	left := 0
	right := len(nums) -1
	end := right
	for left <= right {
		mid := left + (right-left) >> 1
		if nums[mid] <= target {
			left = mid + 1
		}else if nums[mid] > target {
			right = mid - 1
		}
	}

	if right == end {
		return -1
	}

	return right + 1
}

//获取第一个大于等于的数
func findFirstGe2(nums []int,target int) int  {
	left := 0
	right := len(nums) -1
	end := right
	for left <= right {
		mid := left + (right-left) >> 1
		if nums[mid] < target {
			left = mid + 1
		}else if nums[mid] >= target {
			right = mid - 1
		}
	}

	if right == end {
		return -1
	}

	return right + 1
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
//
//func main() {
//	a := []int{1,2,3,4,4,5,6,7}
//	fmt.Println(findFirstGe(a,0,len(a)-1,6))
//}
