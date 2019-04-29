package main

import "fmt"

func numDistinct(s string, t string) int {
	sLen := len(s) + 1
	tLen := len(t) + 1
	dp := make([][]int, tLen)
	for i := 0; i < tLen; i++ {
		dp[i] = make([]int, sLen)
	}

	for i := 0; i < sLen; i++ {
		dp[0][i] = 1
	}

	for i := 1; i < tLen; i++ {
		dp[i][0] = 0
	}
	for i := 1; i < tLen; i++ {
		for j := 1; j < sLen; j++ {
			if s[j-1] == t[i-1] {
				dp[i][j] = dp[i][j-1] + dp[i-1][j-1]
			} else {
				dp[i][j] = dp[i][j-1]
			}
		}
	}

	return dp[tLen-1][sLen-1]
}
func main() {

	fmt.Println(numDistinct("rab", "rabbit"))
	fmt.Println("vim-go")
}
