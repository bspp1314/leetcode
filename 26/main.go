package main

import "fmt"

func removeDuplicates(nums []int) int {
	n := len(nums)
	if n  <= 1 {
		return n
	}

	left := 0
	right := 1

	for right < n {
		 if nums[left] == nums[right] {
		 	right++
		 }else{
		 	nums[left+1] = nums[right]
		 	left++
		 	right++
		 }
	}

	return left + 1
}
func main() {
	nums := []int{1,1,2,2,3,4,4,5,6}
	n := removeDuplicates(nums)
	fmt.Println(nums)
	fmt.Println(n)
}
