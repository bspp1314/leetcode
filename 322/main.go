package main

import (
	"fmt"
	"math"
)


func coinChange(coins []int, amount int) int {

	 exist := map[int]int{}
	for i := 0; i < len(coins); i++ {
		exist[coins[i]] =1
	}

	return help(coins,amount,exist)



}

func help(coins []int, amount int,exist map[int]int) int  {
	if v, ok := exist[amount]; ok {
		return v
	}

	min := math.MaxInt32
	var k int

	for i := 0; i < len(coins); i++ {
		if amount < coins[i] {
			k++
			continue
		}

		if amount == coins[i] {
			exist[amount] = 1
			return 1
		}

		v := help(coins, amount-coins[i],exist) + 1
		if v == 0 {
			k++
			continue
		}

		if v < min {
			min = v
		}
	}

	if k == len(coins) {
		min = -1
	}

	exist[amount] = min
	return min
}

func Min(a, b int) int {
	if a < b {
		return a
	}

	return b
}

func coinChange2(coins []int, amount int) int {
	if amount == 0 {
		return 0
	}
	dp := make([]int, amount+1)

	for i := 1; i < amount; i++ {
		for j := 0; j < len(coins); j++ {
			if coins[i] <= i {
				dp[i] = Min(dp[i], dp[i-coins[j]]+1)
			}
		}
	}

	if dp[amount] > amount {
		return -1
	}

	return dp[amount]

}

func main() {
	fmt.Println(coinChange([]int{2}, 3))
}
