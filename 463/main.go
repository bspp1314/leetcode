package main

import (
	"fmt"
)


var dir4 = [4][2]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}



func islandPerimeter(grid [][]int) (ans int) {
	n, m := len(grid), len(grid[0])
	for i, row := range grid {
		for j, v := range row {
			if v == 1 {
				for _, d := range dir4 {
					if x, y := i+d[0], j+d[1]; x < 0 || x >= n || y < 0 || y >= m || grid[x][y] == 0 {
						ans++
					}
				}
			}
		}
	}
	return
}



func islandPerimeter2(grid [][]int) (ans int) {
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if grid[i][j] == 1 {
				level := [][2]int{[2]int{i, j}}

				cnd := 0
				for len(level) != 0 {
					newLevel := make([][2]int, 0)
					for l := 0; l < len(level); l++ {
						if grid[level[l][0]][level[l][1]] == 2 {
							continue
						} else {
							grid[level[l][0]][level[l][1]] = 2
						}

						for k := 0; k < 4; k++ {
							x := dir4[k][0] + level[l][0]
							y := dir4[k][1] + level[l][1]
							if x < 0 || y < 0 || x >= len(grid) || y >= len(grid[0]) || grid[x][y] == 0 {
								cnd++
							} else {
								newLevel = append(newLevel, [2]int{x, y})
							}
						}
					}
					level = newLevel
				}
				if cnd > ans {
					ans = cnd
				}
			}
		}
	}

	return ans
}

func islandPerimeter3(grid [][]int)  (ans int) {
	max := 0
	var dfs func(x,y int)
	dfs = func(x, y int) {
		if x < 0 || y < 0 || x >= len(grid) || y >= len(grid[0]) || grid[x][y] == 0 {
			ans++
			return
		}

		if grid[x][y] == 2 {
			return
		}

		grid[x][y] = 2

		dfs(x+1, y)
		dfs(x-1, y)
		dfs(x, y+1)
		dfs(x, y-1)
		return
	}

	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if grid[i][j] == 1 {
					dfs(i,j)
			}
		}
	}
	return max
}


func islandPerimeter4(grid [][]int)  (ans int) {
	max := 0
	var dfs func(x,y int,cnd *int)
	dfs = func(x, y int,cnd *int) {
		if x < 0 || y < 0 || x >= len(grid) || y >= len(grid[0]) || grid[x][y] == 0 {
			*cnd++
			return
		}

		if grid[x][y] == 2 {
			return
		}

		grid[x][y] = 2

		dfs(x+1, y,cnd)
		dfs(x-1, y,cnd)
		dfs(x, y+1,cnd)
		dfs(x, y-1,cnd)
		return
	}

	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if grid[i][j] == 1 {
				cnd := 0
				dfs(i,j,&cnd)
				if cnd > ans {
					ans = cnd
				}
			}
		}
	}
	return max
}



func main() {
	v := [][]int{
		{0, 1, 0, 0},
		{1, 1, 1, 0},
		{0, 1, 0, 0},
		{1, 1, 0, 0},
	}

	out := islandPerimeter(v)
	fmt.Println(out)

	v2 := [][]int{
		{0, 1, 0, 0},
		{1, 1, 1, 0},
		{0, 1, 0, 0},
		{1, 1, 0, 0},
	}

	out2 := islandPerimeter2(v2)
	fmt.Println(out2)


	v3 := [][]int{
		{0, 1, 0, 0},
		{1, 1, 1, 0},
		{0, 1, 0, 0},
		{1, 1, 0, 0},
	}

	out3 := islandPerimeter(v3)
	fmt.Println(out3)

	v4 := [][]int{
		{0, 1, 0, 0},
		{1, 1, 1, 0},
		{0, 1, 0, 0},
		{1, 1, 0, 0},
	}

	out4 := islandPerimeter2(v4)
	fmt.Println(out4)
}

// (1 + 15) * 15 / 2
// 4
