package main

import "fmt"

//暴力法
func uniquePaths(m int, n int) int {
	if m == 0 || n == 0 || (m == 1) && (n == 1) {
		return 0
	}
	if m == 1 || n == 1 {
		return 1
	}
	right := uniquePaths(m-1, n)
	down := uniquePaths(m, n-1)
	return right + down
}

func MemuniquePaths(m int, n int) int {
	dp := make([][]int, m+1)
	for i := 0; i <= m; i++ {
		dp[i] = make([]int, n+1)
		for j := 0; j <= n; j++ {
			dp[i][j] = -1
		}
	}
	return MemuniquePathsHelp(m, n, 1, 1, dp)
}
func MemuniquePathsHelp(m int, n int, i, j int, dp [][]int) int {
	if dp[i][j] >= 0 {
		return dp[i][j]
	}

	if m == 1 || n == 1 {
		return 1
	}
	dp[i][j] = MemuniquePathsHelp(m-1, n, i+1, j, dp) + MemuniquePathsHelp(m, n-1, i, j+1, dp)
	return dp[i][j]
}
func BootomUpLcs(m int, n int) int {
	dp := make([][]int, m+1)
	for i := 0; i <= m; i++ {
		dp[i] = make([]int, n+1)
		for j := 0; j <= n; j++ {
			if i == 1 && j > 1 {
				dp[i][j] = 1
			} else if j == 1 && i > 1 {
				dp[i][j] = 1
			} else {
				dp[i][j] = 0
			}
		}
	}
	for i := 2; i <= m; i++ {
		for j := 2; j <= n; j++ {
			dp[i][j] = dp[i-1][j] + dp[i][j-1]
		}
	}
	return dp[m][n]
}
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
func uniquesPathsMath(m, n int) int {
	if m == 0 || n == 0 {
		return 0
	}
	N := m + n - 2
	M := min(m-1, n-1)
	ans := 1
	for i := 1; i <= M; i++ {
		ans = ans * (N - i + 1) / i
	}
	return ans
}
func main() {
	//fmt.Println(uniquePaths(7, 3))
	//fmt.Println(MemuniquePaths(100, 10))
	//fmt.Println(BootomUpLcs(10, 100))
	fmt.Println((uniquesPathsMath(1, 1)))
	fmt.Println((uniquesPathsMath(7, 3)))
	fmt.Println((uniquesPathsMath(10, 100)))
}
