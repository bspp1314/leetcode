package main

import "fmt"

//暴力法
func twoSum(nums []int, target int) []int {
	l := len(nums)
	result := make([]int, 2, 2)
	for i := 0; i < l; i++ {
		for j := i + 1; j < l; j++ {
			if (nums[i] + nums[j]) == target {
				result[0] = i
				result[1] = j
				return result
			}
		}
	}
	return result
}

//Hash
func twoSum2(nums []int, target int) []int {
	l := len(nums)
	result := make([]int, 2, 2)
	tmp := make(map[int]int)

	for i := 0; i < l; i++ {
		if j, ok := tmp[target-nums[i]]; ok {
			result[0] = j
			result[1] = i
			return result
		} else {
			tmp[nums[i]] = i
		}
	}
	return result
}

func main() {
	nums := []int{1, 2, 3, 4, 5, 6}
	result := twoSum(nums, 11)
	fmt.Println(result)
	result2 := twoSum2(nums, 11)
	fmt.Println(result2)
}
