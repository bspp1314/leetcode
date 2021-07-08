package main

import "fmt"

func findMin(nums []int) int {

	if len(nums) == 1 {
		return nums[0]
	}

	left := 0
	right := len(nums) - 1

	if nums[right] > nums[left] {
		return nums[left]
	}

	for left < right {
		mid := left + (right-left) /2

		if nums[mid] <= nums[right] {
			// mid 可能是最小值
			right = mid
		}else{
			left  = mid +1
		}
	}

	return nums[left]
}


func main() {
	fmt.Println(findMin([]int{3,4,5,1,2}))
	fmt.Println(findMin([]int{4,5,6,7,0,1,2}))
	fmt.Println(findMin([]int{11,13,15,17}))
}
