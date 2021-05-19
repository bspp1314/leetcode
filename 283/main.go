package main

import "fmt"

//给定一个数组 nums，编写一个函数将所有 0 移动到数组的末尾，同时保持非零元素的相对顺序。
func moveZeroes(nums []int) {
	if len(nums) <= 1 {
		return
	}

	left := -1
	right := -1
	// [0,1,2,3]

	for i := 0; i < len(nums); i++ {
		if nums[i] == 0 {
			right = i

			if left == -1 {
				left = i
			}
		}else{
			if left == -1 {
				continue
			}

			nums[left],nums[i] = nums[i],nums[left]
			left++
			right++
		}
	}

}
func main() {
	a := []int{0,1,2,3,0,0,5}
	moveZeroes(a)
	fmt.Println(a)

	b := []int{0,1,2,3,0,0,5}
	moveZeroes(b)
	fmt.Println(b)

}
