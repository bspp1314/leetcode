package main

import (
	"fmt"
)

func maxProfit(k int, prices []int) int {
	if len(prices) <= 1 {
		return 0
	}

	if k > len(prices) / 2 {
		return func() int {
			res := 0
			for i:= 1;i<len(prices);i++ {
				if prices[i] > prices[i-1] {
					res += prices[i] - prices[i-1]
				}
			}

			return res
		}()
	}


	K := k * 2
	maxProfit := make([]int,K)

	for i:=0;i < K ;i += 2  {
		maxProfit[i] = 0 - prices[0]
	}

	max := func(a,b int) int {
		if a > b {
			return a
		}

		return b
	}

	for i:=1;i<len(prices);i++ {
		for j := 0;j < K;j += 2   {
			if j == 0 {
				maxProfit[j] = max(maxProfit[j], 0 - prices[i])
				maxProfit[j+1] = max(maxProfit[j+1], maxProfit[j] + prices[i])
				continue
			}

			maxProfit[j] = max(maxProfit[j],maxProfit[j-1] - prices[i])
			maxProfit[j+1] = max(maxProfit[j+1],maxProfit[j] + prices[i])

		}
	}

	return maxProfit[len(maxProfit)-1]
}

func main() {
	out := maxProfit(2,[]int{2,4,1})
	fmt.Println(out)
}
