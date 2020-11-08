package main

import (
	"fmt"
)

func uniquePaths(m int, n int) int {
	dp := make([][]int,m)
	for i:=0;i<len(dp);i++ {
		dp[i] = make([]int,n)
	}


	for i := 0;i < m;i++ {
		for j:=0;j< n;j++ {
			if i == 0 || j == 0 {
				dp[i][j] = 1
			}else{
				dp[i][j] = dp[i-1][j] + dp[i][j-1]
			}
		}
	}

	return dp[m-1][n-1]
}

func uniquePaths2(m int, n int) int {
	dp := make([]int,n)

	for i := 0;i < m;i++ {
		for j:=0;j< n;j++ {
			if i == 0 || j == 0 {
				dp[j] = 1
			}else{
				//当前 dp[j] 保存上方的值
				res  := dp[j]+dp[j-1]
				dp[j] = res
			}
		}
	}

	return dp[n-1]
}

func main() {
	out := uniquePaths(7,3)
	fmt.Println(out)
	out = uniquePaths2(7,3)
	fmt.Println(out)
}
