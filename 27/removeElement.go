package main

import "fmt"

func removeElement(nums []int, val int) int {
	j := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] != val {
			nums[j] = nums[i]
			j++
		}
	}
	return j
}
func main() {
	nums := []int{0, 1, 2, 3, 4, 4, 4, 5, 6}
	index := removeElement(nums, 7)
	fmt.Println(index, "nums", nums)
}
