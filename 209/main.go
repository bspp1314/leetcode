package main

import "fmt"

func minSubArrayLen(s int, nums []int) int {
	if len(nums) == 0 {
		return  0
	}

	left := 0
	right := 0
	sum := 0
	res := 1 << 63 - 1

	for right < len(nums) {
		sum += nums[right]

		for right < left && sum > s {
			if res > left-right {
				res = left -right
			}

			sum = sum - nums[right]
			right++
		}

		left++
	}

	return sum
}

func main() {
	fmt.Println(minSubArrayLen(7,[]int{1,2,4,3}))
}