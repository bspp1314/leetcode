package main

import "fmt"

// 1->6->5->4->3->3->2
func findDuplicate(nums []int) int {
	if len(nums) <= 1 {
		return 0
	}

	if len(nums) == 2 {
		return nums[0]
	}

	slow := 0
	fast := 0

	// flow 算法。
	for  {
		slow = nums[slow]
		fast = nums[nums[fast]]
		if slow == fast {
			break
		}
	}

	slow = 0

	for  {
		slow = nums[slow]
		fast = nums[fast]

		if slow == fast {
			break
		}
	}

	return slow
}
func main() {
	fmt.Println(findDuplicate([]int{1,2,3,3,4,5}))
}
