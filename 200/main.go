package main

import "fmt"

func numIslands2(grid [][]byte) int {
	if len(grid) == 0 || len(grid[0]) == 0 {
		return 0
	}

	res := 0
	var queue [][2]int

	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if grid[i][j] == '0' {
				continue
			}

			queue = [][2]int{[2]int{i, j}}

			for len(queue) != 0 {
				node := queue[0]
				newI, newJ := node[0], node[1]
				queue = queue[1:]
				if newI < 0 || newJ < 0 || newI >= len(grid) || newJ >= len(grid[0]) || grid[newI][newJ] == '0' {
					continue
				}

				grid[newI][newJ] = '0'
				queue = append(queue, [2]int{newI + 1, newJ})
				queue = append(queue, [2]int{newI - 1, newJ})
				queue = append(queue, [2]int{newI, newJ + 1})
				queue = append(queue, [2]int{newI, newJ - 1})
			}

			res++
		}
	}

	return res
}

func numIslands(grid [][]byte) int {
	if len(grid) == 0 || len(grid[0]) == 0 {
		return 0
	}

	res := 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if grid[i][j] == '1' {
				res++
				dfs(grid, i, j)
			}
		}
	}

	return res
}

func dfs(grid [][]byte, i, j int) {
	if i < 0 || j < 0 || i >= len(grid) || j >= len(grid[0]) {
		return
	}

	if grid[i][j] == '0' {
		return
	}

	grid[i][j] = '0'

	dfs(grid, i+1, j)
	dfs(grid, i-1, j)
	dfs(grid, i, j-1)
	dfs(grid, i, j+1)
}

func main() {
	land := [][]byte{
		{'1', '1', '0', '1'},
		{'1', '1', '0', '1'},
		{'1', '1', '0', '1'},
		{'1', '1', '0', '1'},
	}

	out := numIslands2(land)

	fmt.Println(out)
}
