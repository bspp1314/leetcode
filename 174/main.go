package main

import (
	"fmt"
	"math"
)

func Min(a, b int) int {
	if a < b {
		return a
	}

	return b
}

func Max(a, b int) int {
	if a > b {
		return a
	}

	return b
}

func calculateMinimumHP(dungeon [][]int) int {
	n, m := len(dungeon), len(dungeon[0])

	if m == 0 || n == 0 {
		return 0
	}

	dp := make([][]int,n+1)
	for i := 0; i < n+1 ; i++ {
		dp[i] = make([]int,m+1)
		for j := 0; j < m+1; j++ {
			dp[i][j] = math.MaxInt32
		}
	}

	dp[n][m-1] = 1
	dp[n-1][m] = 1

	for i := n-1; i >=0 ; i-- {
		for j := m-1; j >=0 ; j++ {
			min := Min(dp[i+1][j],dp[i][j+1])
			dp[i][j] = Max(min-dungeon[i][j],1)
		}
	}


	return dp[0][0]
}

func main() {
	zz := calculateMinimumHP([][]int{
		{0, -5},
		{-5, 0},
	})

	fmt.Println(zz)
}
