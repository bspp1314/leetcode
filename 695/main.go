package main

import "fmt"

func maxAreaOfIsland2(grid [][]int) int {
	max := 0

	var stack [][2]int
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if grid[i][j] == 0 {
				continue
			}
			ans := 0
			stack = append(stack, [2]int{i, j})
			for len(stack) != 0 {
				node := stack[0]
				stack = stack[1:]
				if node[0] < 0 || node[1] < 0 || node[0] >= len(grid) || node[1] >= len(grid[0]) ||
					grid[node[0]][node[1]] == 0 {
					continue
				}

				ans += 1
				grid[node[0]][node[1]] = 0
				stack = append(stack, [2]int{node[0] + 1, node[1]})
				stack = append(stack, [2]int{node[0] - 1, node[1]})
				stack = append(stack, [2]int{node[0], node[1] + 1})
				stack = append(stack, [2]int{node[0], node[1] - 1})
			}

			if ans > max {
				max = ans
			}
		}
	}

	return max
}

func maxAreaOfIsland(grid [][]int) int {
	max := 0

	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			v := dfs(grid, i, j)
			if v > max {
				max = v
			}
		}
	}

	return max
}

func dfs(grid [][]int, i int, j int) int {
	if i < 0 || j < 0 || i >= len(grid) || j >= len(grid[0]) || grid[i][j] == 0 {
		return 0
	}

	grid[i][j] = 0
	ans := 0
	ans += dfs(grid, i+1, j)
	ans += dfs(grid, i-1, j)
	ans += dfs(grid, i, j+1)
	ans += dfs(grid, i, j-1)
	ans += 1

	return ans
}
func main() {
	v := [][]int{[]int{0, 1}}
	out := maxAreaOfIsland(v)
	fmt.Println(out)
	v2 := [][]int{[]int{0, 1}}
	fmt.Println(maxAreaOfIsland2(v2))
}
