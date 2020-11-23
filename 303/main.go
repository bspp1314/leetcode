package main

import "fmt"

type NumArray struct {
	Dp   []int
	Nums []int
}


func Constructor(nums []int) NumArray {
	dp := make([]int,len(nums))
	if len(nums) == 0 {
		return NumArray{}
	}
	dp[0] = nums[0]
	for i := 1;i<len(nums);i++ {
		dp[i]  = dp[i-1] + nums[i]
	}
	return NumArray{
		Dp:   dp,
		Nums: nums,
	}
}

func (this *NumArray) SumRange(i int, j int) int {
	if len(this.Nums) == 0 {
		return 0
	}

	if i == 0 {
		return this.Dp[j]
	}else{

	}

	return this.Dp[j] - this.Dp[i-1]
}

func main() {
	a := Constructor([]int{-2,0,3,-5,2,-1})
	for i := 0; i < len(a.Dp); i++ {
		fmt.Printf("%d ",a.Dp[i])
	}
	fmt.Println()
	fmt.Println(a.Dp[1]-a.Dp[0])
}
