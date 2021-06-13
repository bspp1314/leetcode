package main

func Max(a, b int) int {
	if a > b {
		return a
	}

	return b
}

//[7,1,5,3,6,4]
func maxProfit(prices []int) int {
	if len(prices) <= 1 {
		return 0
	}

	dp := make([][2]int, len(prices))
	dp[0][0] = 0             //没有持有股票的收益
	dp[0][1] = 0 - prices[0] //持有股票的收益

	for i := 1; i < len(prices); i++ {
		dp[i][0] = Max(dp[i-1][0], dp[i-1][1]+prices[i])
		dp[i][1] = Max(dp[i-1][1], dp[i-1][0]-prices[i])
	}

	return dp[len(prices)-1][0]
}

func maxProfit2(prices []int) int {
	if len(prices) <= 1 {
		return 0
	}

	dp0 := 0             //没有持有股票的收益
	dp1 := 0 - prices[0] //持有股票的收益

	for i := 1; i < len(prices); i++ {
		old := dp0
		dp0 = Max(dp0, dp1+prices[i])
		dp1 = Max(dp1, old-prices[i])
	}

	return dp0
}


func main() {

}
