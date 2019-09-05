package main

import "fmt"

func climbStairs(n int) int {
	if n <= 1 {
		return 1
	}
	dp := make([]int, n+1)
	dp[0] = 1
	for i := 1; i <= n; i++ {
		if i == 1 {
			dp[i] = dp[i-1]
			continue
		}
		dp[i] = dp[i-2] + dp[i-1]
	}
	return dp[n]

}
func main() {
	fmt.Println(climbStairs(10))
	fmt.Println(climbStairs(20))
	fmt.Println("vim-go")
}
