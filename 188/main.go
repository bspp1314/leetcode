package main

func maxProfit(k int, prices []int) int {
	n := len(prices)
	if n == 0 {
		return 0
	}

	k = Min(k, n/2)

	if k == 0 {
		return 0
	}

	buy := make([]int, k)
	sell := make([]int, k)

	for i := 0; i < k; i++ {
		buy[i] = -prices[i]
	}

	for i := 0; i < len(prices); i++ {
		for j := 0; j < k; j++ {
			if j == 0 {
				buy[j] = Max(buy[j], 0-prices[i])
				sell[j] = Max(sell[j], buy[j]+prices[i])
			} else {
				buy[j] = Max(buy[j], sell[j-1]-prices[i])
				sell[j] = Max(sell[j], buy[j]+prices[i])
			}
		}
	}

	return sell[k-1]

}

func Max(a, b int) int {
	if a > b {
		return a
	}

	return b
}

func Min(a, b int) int {
	if a < b {
		return a
	}

	return b
}

func main() {

}
