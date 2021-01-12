package main

import "fmt"

func findErrorNums(nums []int) []int {
	out := make([]int, 2)
	outMap := make(map[int]bool)
	for _, num := range nums {
		if _, ok := outMap[num]; ok {
			out[0] = num
		} else {
			outMap[num] = true
		}

	}
	for i := 0; i < len(nums); i++ {
		if _, ok := outMap[i+1]; !ok {
			out[1] = i + 1
		}
	}
	return out
}

func main() {
	fmt.Println(findErrorNums([]int{1,3,3,4,5}))
}
