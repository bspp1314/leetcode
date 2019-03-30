package main

import "fmt"

func removeDuplicates(nums []int) int {
	l := len(nums)
	if l == 0 {
		return 0
	}
	index := 1
	for i := 1; i < l; i++ {
		if nums[i] != nums[i-1] {
			fmt.Println(i)
			nums[index] = nums[i]
			index++
		}
	}
	return index
}
func main() {
	nums := []int{0, 0, 0, 0, 0, 0, 0}
	removeDuplicates(nums)
	fmt.Println(nums)
}
