package main

import "fmt"

func missingNumber(nums []int) int {
	n := len(nums)
	v := ((1+n) * n) /2

	for i := 0; i < len(nums); i++ {
		v -= nums[i]
	}

	return v
}
func main()  {
	fmt.Println(missingNumber([]int{1,0,3}))
}
