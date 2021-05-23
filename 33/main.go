package main

import "fmt"

//[4,5,6,7,0,1,2], target =
func searchK(nums []int, target int) int {
	if len(nums) == 0 {
		return -1
	}

	left := 0
	right := len(nums) - 1
	n := len(nums)

	for left <= right {
		mid := left + (right-left)/2
		if nums[mid] == target {
			return mid
		}

		// mid值在左边
		if nums[0] <= nums[mid] {
			if target >= nums[0] && target < nums[mid] {
				right = mid - 1
			} else {
				left = left + 1
			}
		}else{
			if nums[mid] < target && target <= nums[n-1] {
				left = mid + 1
			}else{
				right = mid -1
			}
		}

	}

	return -1

}

func main() {
	nums := []int{7, 8, 9, 10, 1, 2, 3, 4, 5, 6}
		fmt.Println(searchK(nums, 1))
	fmt.Println(searchK(nums, 2))
	fmt.Println(searchK(nums, 3))
	fmt.Println(searchK(nums, 4))
	fmt.Println(searchK(nums, 5))
	fmt.Println(searchK(nums, 6))
	fmt.Println(searchK(nums, 7))
	fmt.Println(searchK(nums, 8))
	fmt.Println(searchK(nums, 9))
	fmt.Println(searchK(nums, 10))
}
