package main

import "fmt"

func maxProfit(prices []int) int {
	dp := make([][3]int,len(prices))
	dp[0][0] =  0              //sell max(dp[i-1][0],dp[i] + prices[i])
	dp[0][1] =  0 - prices[0] // buy  max(dp[i-1][1],dp[i-1][2] - prices[i])
	dp[0][2] =  0      // block

	max := func(a,b int) int {
		if a > b {
			return a
		}

		return b
	}

	for i:=1;i<len(prices);i++ {
		dp[i][0] = max(dp[i-1][0],dp[i-1][1] + prices[i])  // sell
		dp[i][1] = max(dp[i-1][1],dp[i-1][2] - prices[i]) // buy
		dp[i][2] = dp[i-1][0]
	}

	return dp[len(prices)-1][0]
}
func main() {
	fmt.Println(maxProfit([]int{1,2,3,0,2}))
}
