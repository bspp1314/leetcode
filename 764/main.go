package main

import (
	"fmt"
)

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
	for k := 0; k < N; k++ { //k代表第k行，第k列
		left := 0
		right := 0
		up := 0
		down := 0 //分别代表上下左右四个方向连续的连续非0的个数
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
	out := orderOfLargestPlusSign(5, nil)
	fmt.Println("out is ", out)
}
