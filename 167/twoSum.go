package main

import "fmt"

//Hash
func twoSum(nums []int, target int) []int {
	l := len(nums)
	result := make([]int, 2, 2)
	tmp := make(map[int]int)

	for i := 0; i < l; i++ {
		if j, ok := tmp[target-nums[i]]; ok {
			result[0] = j + 1
			result[1] = i + 1
			return result
		} else {
			tmp[nums[i]] = i
		}
	}
	return result
}

//双指针法
func twoSum2(nums []int, target int) []int {
	p1 := 0
	p2 := len(nums) - 1
	result := make([]int, 2, 2)

	for p1 != p2 {
		p := nums[p1] + nums[p2]
		if p > target {
			p2--
		} else if p < target {
			p1++
		} else {
			result[0] = p1 + 1
			result[1] = p2 + 1
			return result
		}
	}

	return result
}

func main() {
	var q float64
	fmt.Printf("%0.0f\n", q)
	nums := []int{1, 2, 3, 4, 5, 6}
	fmt.Println(twoSum(nums, 11))
	fmt.Println(twoSum2(nums, 11))
}
