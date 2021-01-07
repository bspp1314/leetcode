package main

import (
	"fmt"
)

func Min(a,b int) int  {
	if a < b {
		return a
	}
	return b
}

func orderOfLargestPlusSign2(N int, mines [][]int) int {
	if N == 0 {
		return 0
	}

	dp := make([][]int, N)
	for i := 0; i < N; i++ {
		dp[i] = make([]int, N)
		for j := 0; j < N; j++ {
			dp[i][j] = N
		}
	}

	for i := 0; i < len(mines); i++ {
		dp[mines[i][0]][mines[i][1]] = 0
	}


	for i := 0; i < N; i++ {
		left := 0
		right := 0
		down := 0
		up := 0
		for j := 0; j < N; j++ {
			//left
			if dp[i][j] == 0 {
				left = 0
			}else{
				left = left +  1
			}

			dp[i][j] = Min(dp[i][j],left)

			//right
			if dp[i][N-j-1] == 0 {
				right = 0
			}else{
				right = right +1
			}
			dp[i][N-j-1] = Min(dp[i][N-j-1],right)

			// down
			if dp[j][i] == 0 {
				down = 0
			}else{
				down = down + 1
			}

			dp[j][i] = Min(dp[j][i],down)

			//up
			if dp[N-j-1][i] == 0 {
				up = 0
			}else{
				up = up+1
			}

			dp[N-j-1][i] = Min(dp[N-j-1][i],up)
		}
	}

	max := 0
	for i := 0; i < N; i++ {
		for j := 0; j < N; j++ {
			if dp[i][j] > max {
				max = dp[i][j]
			}

		}
	}

	return max
}

func orderOfLargestPlusSign(N int, mines [][]int) int {
	if N == 0 {
		return 0
	}

	min := func(a, b int) int {
		if a < b {
			return a
		}
		return b
	}

	grid := make([][]int, N)
	for i := 0; i < N; i++ {
		grid[i] = make([]int, N)
		for j := 0; j < N; j++ {
			grid[i][j] = N
		}
	}

	for i := 0; i < len(mines); i++ {
		grid[mines[i][0]][mines[i][1]] = 0
	}

	// k 表示第 k 行 第 k 列
	for k := 0; k < N; k++ {
		left := 0
		right := 0
		up := 0
		down := 0
		for i, j :=0, N - 1; i < N; i, j = i+1, j-1 {
			//第k行的左边扫描
			if grid[k][i] == 0 {
				left = 0
			} else {
				left = left + 1
			}
			grid[k][i] = min(grid[k][i], left)

			if grid[k][j] == 0 {
				right = 0
			} else {
				right += 1
			}

			grid[k][j] = min(grid[k][j], right)

			if grid[i][k] == 0 {
				up = 0
			} else {
				up += 1
			}
			grid[i][k] = min(grid[i][k], up)

			if grid[j][k] == 0 {
				down = 0
			} else {
				down += 1
			}
			grid[j][k] = min(grid[j][k], down)

		}
	}

	max := 0
	for i := 0; i < N; i++ {
		for j := 0; j < N; j++ {
			if grid[i][j] > max {
				max = grid[i][j]
			}
		}
	}

	return max
}

func main() {
	fmt.Println(orderOfLargestPlusSign(5, [][]int{[]int{4,2}}))
	fmt.Println(orderOfLargestPlusSign2(5, [][]int{[]int{4,2}}))

}
