package main

import "fmt"

func uniquePathsWithObstacles2(obstacleGrid [][]int) int {
	m := len(obstacleGrid)
	if m == 0 {
		return 0
	}
	n := len(obstacleGrid[0])
	if n == 0 {
		return 0
	}
	if obstacleGrid[0][0] == 1 {
		return 0
	}
	obstacleGrid[0][0] = 1
	for i := 1; i < m; i++ {
		if obstacleGrid[i][0] == 0 && obstacleGrid[i-1][0] == 1 {
			obstacleGrid[i][0] = 1
		} else {
			obstacleGrid[i][0] = 0
		}
	}

	for j := 1; j < n; j++ {
		if obstacleGrid[0][j] == 0 && obstacleGrid[0][j-1] == 1 {
			obstacleGrid[0][j] = 1
		} else {
			obstacleGrid[0][j] = 0
		}
	}

	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			if obstacleGrid[i][j] == 0 {
				obstacleGrid[i][j] = obstacleGrid[i-1][j] + obstacleGrid[i][j-1]
			} else {
				obstacleGrid[i][j] = 0
			}
		}
	}
	return obstacleGrid[m-1][n-1]

}
func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	m := len(obstacleGrid)
	if m == 0 {
		return 0
	}
	n := len(obstacleGrid[0])
	if n == 0 {
		return 0
	}
	dp := make([][]int, m+1)
	for i := 1; i <= m; i++ {
		dp[i] = make([]int, n+1)
	}
	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if obstacleGrid[i-1][j-1] == 1 {
				continue
			}
			if i == 1 && j == 1 {
				dp[i][j] = 1
			} else if i == 1 {
				if obstacleGrid[i-1][j-2] == 0 {
					dp[i][j] = dp[i][j-1]
				}
			} else if j == 1 {
				if obstacleGrid[i-2][j-1] == 0 {
					dp[i][j] = dp[i-1][j]
				}
			} else {
				if obstacleGrid[i-2][j-1] == 0 {
					dp[i][j] += dp[i-1][j]
				}
				if obstacleGrid[i-1][j-2] == 0 {
					dp[i][j] += dp[i][j-1]
				}
			}
		}
	}
	return dp[m][n]
}
func main() {
	in := [][]int{
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0},
		{0, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 1},
	}
	fmt.Println(uniquePathsWithObstacles(in))

	in = [][]int{{0, 1}}
	fmt.Println(uniquePathsWithObstacles(in))
}
