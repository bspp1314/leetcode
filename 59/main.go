package main

import "fmt"

func generateMatrixHelp(matrix [][]int, left int, right int, top int, down int, key *int) {
	if left > right {
		return
	}

	if top > down {
		return
	}

	if left == right {
		for i := top; i <= down; i++ {
			matrix[i][right] =  *key
			*key = *key + 1
		}
		return
	}

	if top == down {
		for i := left; i <= right; i++ {
			matrix[top][i] = *key
			*key = *key + 1
		}
		return
	}

	for i := left; i < right; i++ {
		matrix[top][i] = *key
		*key = *key + 1
	}

	for i := top; i < down; i++ {
		matrix[i][right] = *key
		*key = *key + 1
	}

	for i := right; i > left; i-- {
		matrix[down][i] = *key
		*key = *key + 1
	}

	for i := down; i > top; i-- {
		matrix[i][left] = *key
		*key = *key + 1
	}

	generateMatrixHelp(matrix, left+1, right-1, top+1, down-1, key)
}

func generateMatrix(n int) [][]int {
	matrix := make([][]int,n)
	for i := 0; i < n; i++ {
		matrix[i] = make([]int,n)
	}

	key := 1

	top := 0
	down := len(matrix) - 1
	left := 0
	right := len(matrix[0]) - 1

	generateMatrixHelp(matrix,left,right,top,down,&key)

	return matrix
}

func generateMatrix3(n int) [][]int {
	if n == 1 {
		return [][]int{{1}}
	}

	if n == 2 {
		return [][]int{{1,2},{4,3}}
	}

	res := make([][]int,n)
	v := 1

	for i := 0;i < n ;i++ {
		res[i] = make([]int,n)
	}

	left := 0
	right := n-1
	top := 0
	down := n -1

	for left <= right && top <= down {
		if right -left == 0 {
			res[down][left] = v
			break
		}

		if right-left == 1 {
			res[top][left] = v
			res[top][right] = v + 1
			res[down][right] = v +2
			res[down][left] = v + 3
			break
		}

		for i := left;i <= right;i++ {
			res[top][i] = v
			v++
		}

		for i := top+1;i <= down;i++ {
			res[i][right] = v
			v++
		}

		for i := right-1;i > left;i-- {
			res[down][i] = v
			v++
		}

		for i := down;i >= top+1;i-- {
			res[i][left] = v
			v++
		}
		left++
		right--
		top++
		down--

	}

	return res


}
func main()  {
	fmt.Println(generateMatrix3(3))
	//fmt.Println(generateMatrix(5))
}

