package main

import (
	"fmt"
	"math"
)

func findNumbers(nums []int) int {
	res := 0
	for i := 0; i < len(nums); i++ {
		k := int(math.Log10(float64(nums[i])))
		if k%2 == 0 {
			res++
		}
	}

	return res

}
func main() {
	fmt.Println(findNumbers([]int{1, 22, 33, 44, 55, 555, 555, 567}))
}
