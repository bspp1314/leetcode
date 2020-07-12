package main

import "fmt"

func searchRange(nums []int, target int) []int {
	res := []int{-1, -1}
	if len(nums) == 0 {
		return res
	}

	return search(nums, 0, len(nums)-1, target)
}

func search(nums []int, left int, right int, taget int) []int {
	res := []int{-1, -1}

	for left <= right {
		mid := left + (right-left)/2

		if taget == nums[mid] {
			res[0] = mid
			res[1] = mid
			leftRange := search(nums, left, mid-1, taget)
			rightRange := search(nums, mid+1, right, taget)

			if leftRange[0] != -1 {
				res[0] = leftRange[0]
			}

			if rightRange[1] != -1 {
				res[1] = rightRange[1]
			}

			return res
		} else if taget > nums[mid] {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	return res
}

func main() {
	fmt.Println("vim-go")
}
