package main

import (
	"fmt"
)

func Max(a,b int) int   {
	if a > b {
		return a
	}

	return b
}

func Min(a,b int) int   {
	if a < b {
		return a
	}

	return b
}



func maxProduct(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	dpMax := make([]int,len(nums))
	dpMin := make([]int,len(nums))
	dpMax[0] = nums[0]
	dpMin[0] = nums[0]

	res := 0

	for i := 1; i < len(nums); i++ {
		newMax := nums[i] * dpMax[i-1]
		newMin := nums[i] * dpMin[i-1]
		if newMax < newMin {
			newMax,newMin = newMin,newMax
		}

		dpMax[i] = Max(newMax,nums[i])
		dpMin[i] = Min(newMin,nums[i])

		res = Max(res,dpMax[i])
	}

	return res
}

func maxProduct2(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	product := nums[0]
	max := nums[0]
	min := nums[0]

	for i := 1; i < len(nums); i++ {
		newMax := nums[i] * max
		newMin := nums[i] * min
		if newMax < newMin {
			newMax,newMin = newMin,newMax
		}

		min = Min(nums[i],newMin)
		max = Max(nums[i],newMax)
		product = Max(max,product)

	}

	return product
}



func main() {
	o := maxProduct([]int{-2,3,-4})
	o2 := maxProduct2([]int{-2,3,-4})
	fmt.Println(o)
	fmt.Println(o2)
}