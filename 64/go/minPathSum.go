package main

import "fmt"

func minPathSum1(grid [][]int) int {
	m := len(grid)
	if m == 0 {
		return 0
	}
	n := len(grid[0])
	if n == 0 {
		return 0
	}
	dp := make([][]int, m+1)
	for i := 0; i <= m; i++ {
		dp[i] = make([]int, n+1)
	}
	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if i == 1 {
				dp[i][j] = dp[i][j-1]
			} else if j == 1 {
				dp[i][j] = dp[i-1][j]
			} else {
				if dp[i-1][j] < dp[i][j-1] {
					dp[i][j] = dp[i-1][j]
				} else {
					dp[i][j] = dp[i][j-1]
				}
			}
			dp[i][j] += grid[i-1][j-1]
		}
	}
	return dp[m][n]
}
func minPathSum2(grid [][]int) int {
	m := len(grid)
	if m == 0 {
		return 0
	}
	n := len(grid[0])
	if n == 0 {
		return 0
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if i == 0 && j == 0 {
				continue
			} else if i == 0 {
				grid[i][j] += grid[i][j-1]
			} else if j == 0 {
				grid[i][j] += grid[i-1][j]
			} else {
				if grid[i-1][j] < grid[i][j-1] {
					grid[i][j] += grid[i-1][j]
				} else {
					grid[i][j] += grid[i][j-1]
				}
			}
		}
	}
	return grid[m-1][n-1]
}
func main() {
	in := [][]int{
		{3, 3, 5, 0},
		{1, 2, 3, 0},
		{0, 1, 0, 0},
	}
	in2 := [][]int{
		{3},
		{1},
		{0},
	}
	in3 := [][]int{
		{3, 3, 5, 0},
	}
	in4 := [][]int{
		{3},
	}
	fmt.Println(minPathSum(in))
	fmt.Println(minPathSum(in2))
	fmt.Println(minPathSum(in3))
	fmt.Println(minPathSum(in4))

}
