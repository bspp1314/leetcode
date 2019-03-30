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
	in := []int{1, 9, 3, 4, 5, 6, 7, 8, 9}
	fmt.Println(in)
	fmt.Println(findErrorNums(in))
	in2 := []int{1, 2, 3, 4, 5, 6, 7, 8, 2}
	fmt.Println(in2)
	fmt.Println(findErrorNums(in2))
	in3 := []int{3, 2, 2}
	fmt.Println(findErrorNums(in3))
}
