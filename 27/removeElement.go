package main

import "fmt"

func removeElement(nums []int, val int) int {
	index := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] != val {
			nums[index] = nums[i]
			index++
		}
	}
	return index
}
func main() {
	nums := []int{0, 1, 2, 3, 4, 4, 4, 5, 6}
	index := removeElement(nums, 7)
	fmt.Println(index, "nums", nums)
}
