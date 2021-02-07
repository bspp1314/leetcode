package main

import "fmt"

// get max

// 1
func PredictTheWinner2(nums []int) bool {
	n := len(nums)
	sum := 0
	if n&1 == 0 {
		return true
	}

	dp := make([][]int, n)
	for i := 0; i < len(nums); i++ {
		sum += nums[i]
		dp[i] = make([]int, n)
		for j := 0; j < n; j++ {
			dp[i][j] = -1
		}

		dp[i][i] = nums[i]
	}

	return sum-2*help2(nums, 0, n-1, dp) <= 0
}

func PredictTheWinner(nums []int) bool {
	n := len(nums)
	sum := 0
	if n&1 == 0 {
		return true
	}

	for i := 0; i < len(nums); i++ {
		sum += nums[i]
	}

	return sum-2*help(nums, 0, n-1) <= 0

}

func Min(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}

func Max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func help(nums []int, i, j int) int {
	if i == j {
		return nums[i]
	}

	return Max(
		//  对手在这把 nums[i+1] 或 num[j],让自己有最优解，也就让我方有最差解
		nums[i]+Min(help(nums, i+2, j), help(nums, i+1, j-1)),
		// 对手在这把选 nums[i] 或 num[j-1],让自己有最优解，也就让我方有最差解
		nums[j]+Min(help(nums, i+1, j-1), help(nums, i, j-2)),
	)
}

func help2(nums []int, i, j int, dp [][]int) int {
	if i == j {
		return dp[i][j]
	}
	if dp[i][j] > 0 {
		return dp[i][j]
	}

	dp[i][j] = Max(
		//  对手在这把 nums[i+1] 或 num[j],让自己有最优解，也就让我方有最差解
		nums[i]+Min(help2(nums, i+2, j, dp), help2(nums, i+1, j-1, dp)),
		// 对手在这把选 nums[i] 或 num[j-1],让自己有最优解，也就让我方有最差解
		nums[j]+Min(help2(nums, i+1, j-1, dp), help2(nums, i, j-2, dp)),
	)

	return dp[i][j]
}

func PredictTheWinner3(nums []int) bool {
	n := len(nums)
	if n&1 == 0 {
		return true
	}

	dp := make([][]int, n)
	sum := 0
	for i := 0; i < n; i++ {
		dp[i] = make([]int, n)
		sum += nums[i]
		dp[i][i] = nums[i]
	}

	for j := 1; j < n; j++ {
		dp[j-1][j] = Max(dp[j-1][j-1], dp[j][j])
	}

	for i := n - 3; i >= 0; i-- {
		for j := i + 2; j < n; j++ {
			dp[i][j] = Max(
				nums[i]+Min(dp[i+2][j], dp[i+1][j-1]),
				nums[j]+Min(dp[i+1][j-1], dp[i][j-2]),
			)
		}
	}

	return sum-2*dp[0][n-1] <= 0
}

func main() {
	fmt.Println(PredictTheWinner([]int{7,1, 5, 2,0}))
	fmt.Println(PredictTheWinner2([]int{7,1, 5, 2,0}))
	fmt.Println(PredictTheWinner3([]int{7,1, 5, 2,0}))
}
