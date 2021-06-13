package main

import "math"

func maxProfit(prices []int) int {
	if len(prices) <= 1 {
		return 0
	}

	minPrice := math.MaxInt64
	maxProfitValue := 0
	for i := 0; i < len(prices); i++ {
		if prices[i] < minPrice {
			minPrice = prices[i]
		}else {
			v := prices[i] - minPrice
			if v > maxProfitValue {
				maxProfitValue = v
			}
		}
	}

	return maxProfitValue
}

func main() {

}
