package main

import "fmt"

func searchInsert(nums []int, target int) int {
	//获取第一个大于等于target的数
	if len(nums) == 0 {
		return 0
	}

	left := 0
	right := len(nums) -1
	res := -1
	if target < nums[0] {
		return 0
	}

	if target > nums[right] {
		return right+1
	}

	for left <= right {
		mid := left + (right-left) /2
		if nums[mid] >= target {
			res = mid
			right = mid -1
		}else  {
			left = left +1
		}

	}

	return res

}

func main() {
	a := searchInsert([]int{2,4,6,8,8,10},7)
	fmt.Println(a)
}
