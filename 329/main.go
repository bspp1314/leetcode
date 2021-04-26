package main

import "fmt"

var (
	dirs          = [][]int{[]int{-1, 0}, []int{1, 0}, []int{0, -1}, []int{0, 1}}
	rows, columns int
)

func longestIncreasingPath(matrix [][]int) int {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return 0
	}
	rows, columns = len(matrix), len(matrix[0])
	dp := make([][]int, rows)
	for i := 0; i < rows; i++ {
		dp[i] = make([]int, columns)
	}
	ans := 0
	for i := 0; i < rows; i++ {
		for j := 0; j < columns; j++ {
			ans = max(ans, dfs(matrix, i, j, dp))
		}
	}
	return ans
}

func longestIncreasingPath2(matrix [][]int) int {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return 0
	}
	rows, columns = len(matrix), len(matrix[0])
	dp := make([][]int, rows)
	for i := 0; i < rows; i++ {
		dp[i] = make([]int, columns)
	}
	ans := 0
	for i := 0; i < rows; i++ {
		for j := 0; j < columns; j++ {
			ans = max(ans, dfs(matrix, i, j, dp))
		}
	}
	return ans
}

func dfs(matrix [][]int, row, column int, dp [][]int) int {
	if dp[row][column] != 0 {
		return dp[row][column]
	}
	dp[row][column]++
	for _, dir := range dirs {
		newRow, newColumn := row + dir[0], column + dir[1]
		if newRow >= 0 && newRow < rows && newColumn >= 0 && newColumn < columns && matrix[newRow][newColumn] > matrix[row][column] {			dp[row][column] = max(dp[row][column], dfs(matrix, newRow, newColumn, dp) + 1)
			dp[row][column] = max(dp[row][column],dfs(matrix,newRow,newColumn,dp)+1)
		}
	}
	return dp[row][column]
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func main() {
	out := longestIncreasingPath([][]int{
		{9,9,4},
		{6,6,8},
		{2,1,1},
	})

	fmt.Print(out)
}

