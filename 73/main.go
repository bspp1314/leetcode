package main

import "fmt"

// O(m+n) 的时间复杂度
func setZeroes(matrix [][]int) {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return
	}
	m := len(matrix)
	n := len(matrix[0])
	//记录各行中是否存在0
	rowZero := false
	//记录各列中是否存在0
	colZero := false

	//用第一行和第一列来标识某行某列是否应该被设计为0

	for i := 0; i < m; i++ {
		if matrix[i][0] == 0 {
			colZero = true

		}
	}
	for i := 0; i < n; i++ {
		if matrix[0][i] == 0 {
			rowZero = true
		}
	}

	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			if matrix[i][j] == 0 {
				matrix[i][0] = 0
				matrix[0][j] = 0
				PrintMa(matrix)
			}
		}
	}

	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			if matrix[0][j] == 0 || matrix[i][0] == 0 {
				matrix[i][j] = 0
			}
		}
	}

	//第一行中存在0
	if rowZero {
		for i := 0; i < n; i++ {
			matrix[0][i] = 0
		}
	}

	if colZero {
		for i := 0; i < m; i++ {
			matrix[i][0] = 0
		}
	}

}

func setZeroes2(matrix [][]int) {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return
	}
	m := len(matrix)
	n := len(matrix[0])
	//记录各行中是否存在0
	row := make([]bool, m)
	//记录各列中是否存在0
	col := make([]bool, n)

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if matrix[i][j] == 0 {
				row[i] = true
				col[j] = true
			}
		}
	}

	//遍历行
	for i := 0; i < m; i++ {
		if row[i] {
			for j := 0; j < n; j++ {
				matrix[i][j] = 0
			}
		}
	}

	//遍历列
	for j := 0; j < n; j++ {
		if col[j] {
			for i := 0; i < m; i++ {
				matrix[i][j] = 0
			}
		}
	}
}

func PrintMa(a [][]int) {
	for i := 0; i < len(a); i++ {
		fmt.Println(a[i])
	}
}

func Print(a [][]int, f func(b [][]int)) {
	f(a)
	PrintMa(a)

}

func main() {
	Print([][]int{
		{1, 1, 1,1,1},
		{1, 1, 1,1,1},
		{1, 1, 0,1,1},
		{1, 1, 1,1,1},
		{1, 1, 1,1,1},
	}, setZeroes)
}
