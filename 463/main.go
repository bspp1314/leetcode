package main

import (
	"fmt"
)

func islandPerimeter2(grid [][]int) int {
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
					ans++
					continue
				}
				if grid[node[0]][node[1]] == 2 {
					continue
				}

			}

			if ans > max {
				max = ans
			}
		}
	}

	return max
}


func islandPerimeter(grid [][]int) int {
	max := 0
	for i := 0; i < len(grid); i++ {
		for j:=0;j<len(grid[i]);j++ {
			if grid[i][j] == 1 {
				ans := 0
				dfs(grid,i,j,&ans)
				if ans > max {
					max = ans
				}
			}
		}
	}
	return max
 }

func dfs(grid [][]int, i int, j int,ans *int)  {
	if i < 0 || j < 0 || i >= len(grid) || j >= len(grid[0]) || grid[i][j] == 0 {
		*ans++
		return
	}

	if grid[i][j] == 2 {
		return
	}

	grid[i][j] = 2
	dfs(grid,i+1,j,ans)
	dfs(grid,i-1,j,ans)
	dfs(grid,i,j+1,ans)
	dfs(grid,i,j-1,ans)
	return
}

func main() {
	v := [][]int{
		{0,1,0,0},
		{1,1,1,0},
		{0,1,0,0},
		{1,1,0,1},
	}

	out := islandPerimeter(v)
	fmt.Println(out)

	v2 := [][]int{
		{0,1,0,0},
		{1,1,1,0},
		{0,1,0,0},
		{1,1,0,1},
	}

	out2 := islandPerimeter(v2)
	fmt.Println(out2)
}
