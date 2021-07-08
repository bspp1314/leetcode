package main

import "fmt"
// 偷窃第 k 间房屋，那么就不能偷窃第 k−1 间房屋，偷窃总金额为前 k-2 间房屋的最高总金额与第 k 间房屋的金额之和。

// 不偷窃第 k 间房屋，偷窃总金额为前 k−1 间房屋的最高总金额。


func Max(a,b int) int   {
	if a > b {
		return a
	}

	return b
}

func rob(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	if len(nums) == 1 {
		return nums[0]
	}



	dp0 := nums[0]
	dp1 := Max(dp0,nums[1])

	for i := 1; i < len(nums); i++ {
		new := Max(dp0+nums[i],dp1)
		dp0 = dp1
		dp1 = new

	}

	return dp1
}


func main() {
	out := rob([]int{1,1,3,6,7,10,7,1,8,5,9,1,4,4,3})
	fmt.Println(out)
}
