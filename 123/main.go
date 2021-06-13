package main

import (
	"fmt"
)

func Max(a, b int) int {
	if a > b {
		return a
	}

	return b
}

func maxProfit(prices []int) int {
	if len(prices) <= 1 {
		return 0
	}

	//- 第一买入
	//- 第一次卖出
	//- 第二次买入
	//- 第二次卖出
	dp0 := 0 - prices[0]
	dp1 := 0
	dp2 := 0 - prices[0]
	dp3 := 0

	for i := 0; i < len(prices); i++ {
		dp0 = Max(dp0, -prices[i])
		dp1 = Max(dp1, dp0+prices[i])
		dp2 = Max(dp2, dp1-prices[i])
		dp3 = Max(dp3, dp2+prices[i])
	}

	return dp3
}

func maxProfitK(k int, prices []int) int {
	if len(prices) <= 1 {
		return 0
	}

	if k > len(prices)/2 {
		return func() int {
			res := 0
			for i := 1; i < len(prices); i++ {
				if prices[i] > prices[i-1] {
					res += prices[i] - prices[i-1]
				}
			}

			return res
		}()
	}

	K := k * 2
	maxProfit := make([]int, K)

	for i := 0; i < K; i += 2 {
		maxProfit[i] = 0 - prices[0]
	}

	max := func(a, b int) int {
		if a > b {
			return a
		}

		return b
	}

	for i := 1; i < len(prices); i++ {
		for j := 0; j < K; j += 2 {
			if j == 0 {
				maxProfit[j] = max(maxProfit[j], 0-prices[i])
				maxProfit[j+1] = max(maxProfit[j+1], maxProfit[j]+prices[i])
				continue
			}

			maxProfit[j] = max(maxProfit[j], maxProfit[j-1]-prices[i])
			maxProfit[j+1] = max(maxProfit[j+1], maxProfit[j]+prices[i])

		}
	}

	return maxProfit[len(maxProfit)-1]
}

func main() {
	out := maxProfit([]int{3, 3, 5, 0, 0, 3, 1, 4})
	fmt.Println(out)
}
