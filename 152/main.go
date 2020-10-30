package main

import (
	"fmt"
)


func maxProduct(nums []int) int {
	if len(nums) <= 0 {
		return 0
	}

	maxFun := func(a,b int) int  {
		if a > b {
			return a
		}

		return b
	}

	minFun := func(a,b int) int  {
		if a < b {
			return a
		}

		return b
	}
	//max :=nums[0]
	//min :=nums[0]
	//res := nums[0]
	//
	maxDp := make([]int,0,len(nums))
	minDp := make([]int,0,len(nums))
	maxDp[0] = nums[0]
	minDp[0] = nums[0]

	for i:=1;i<len(nums);i++ {
		//newMax :=maxDp[i-1] * nums[i]
		//newMin := minDp[i-1] * nums[i]
		//if newMax < newMin {
		//	newMin,newMax
		//}
	}

	//for i:=0 ;i <len(nums);i++ {
	//	newMax := max * nums[i]
	//	newMin := min * nums[i]
	//
	//	if newMax < newMin {
	//		newMax,newMin = newMin,newMax
	//	}
	//
	//	min = minFun(newMin,nums[i])
	//	max = maxFun(newMax,nums[i])
	//
	//	if res > max {
	//		res = max
	//	}
	//}

	return max
}



func main() {
	o := maxProduct([]int{-2,3,-4})
	o2 := maxProduct2([]int{-2,3,-4})
	fmt.Println(o)
	fmt.Println(o2)
}