package main

import (
	"fmt"
)

//[2,-1,1,2,2]  0->2->3->0
//[-1,2] 1->1->1->>
// [-10,1,3]  0->3->1->2-1->2

func NextIndex(nums []int, index, l int) int {
	nextStep := nums[index] % l

	if nextStep < 0 {
		nextStep = l + nextStep
	}

	return (nextStep + index) % l

}

func circularArrayLoop(nums []int) bool {
	if len(nums) <= 1 {
		return false
	}
	var res bool
	n := len(nums)

	for i := 0; i < n; i++ {
		slow := i
		fast := NextIndex(nums, i, n)

		//保证方向一致
		for nums[i]*nums[slow] > 0 && nums[i]*nums[fast] > 0 && nums[i]*nums[NextIndex(nums, fast, n)] > 0 {
			if slow == fast {
				//保证环的长度至少为1
				if slow == NextIndex(nums, slow, n) {
					break
				}
				return true
			}

			slow = NextIndex(nums, slow, n)
			fast = NextIndex(nums, NextIndex(nums, fast, n), n)
		}
	}

	return res
}

func main() {
	//fmt.Println(circularArrayLoop([]int{2,-1,1,2,2}))
	//fmt.Println(circularArrayLoop([]int{-1,2}))
	//fmt.Println(circularArrayLoop([]int{-2,1,-1,-2,-2}))
	//fmt.Println(circularArrayLoop([]int{2,-1,1,-2,-2}))
	//fmt.Println(circularArrayLoop([]int{-1,-2,-3,-4,-5}))
	fmt.Println(circularArrayLoop([]int{3, 1, 2}))
	fmt.Println(circularArrayLoop([]int{4, 1, 2, 3}))
	fmt.Println(circularArrayLoop([]int{2,-1,1,-2,-2}))

}
