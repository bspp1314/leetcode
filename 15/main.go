package main

import (
	"fmt"
	"sort"
)

func threeSum(nums []int) [][]int {
	res := make([][]int, 0)
	if len(nums) < 2 {
		return res
	}

	//-1, 0, 1, 2, -1, -4
	//-4  -1 -1 0
	sort.Sort(sort.IntSlice(nums))

	slow := 0
	for nums[slow] <= 0 && (slow+2) < len(nums) {
		if slow > 0 && nums[slow] == nums[slow-1] {
			continue
		}

		left := slow + 1
		right := len(nums) - 1
		for left < right {
			k := nums[left] + nums[right] + nums[slow]
			if k == 0 {
				res = append(res, []int{nums[slow], nums[left], nums[right]})

				for left < right && nums[left] == nums[left+1] { // 去重
					left++
				}

				for left < right && nums[right] == nums[right-1] { // 去重
					right--
				}
				left++
				right--
			} else if k > 0 {
				right--
			} else {
				left++
			}
		}

		slow++
	}

	return res
}

func threeSum2(nums []int) [][]int {
	res := make([][]int, 0)
	if len(nums) < 2 {
		return res
	}

	sort.Ints(nums)

	slow := 0
	for nums[slow] <= 0 && (slow+2) < len(nums) {
		if slow > 0 && nums[slow] == nums[slow-1] {
			continue
		}

		left := slow + 1
		right := len(nums) - 1

		for left < right {
			k := nums[left] + nums[right] + nums[slow]
			if k == 0 {
				res = append(res, []int{slow, left, right})

				for left < right && nums[left] == nums[left+1] {
					left++
				}

				for left < right && nums[right] == nums[right-1] {
					right--
				}
				left++
				right--
			} else if k < 0 {
				left++
			} else {
				right--
			}
		}
	}

	return res
}

func main() {
	fmt.Println(threeSum([]int{0, 0, 0}))
}
