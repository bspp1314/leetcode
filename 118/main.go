package main

import "fmt"

func generate(numRows int) [][]int {
	if numRows == 0 {
		return [][]int{}
	}


	res := make([][]int,numRows)
	res[0] = []int{1}

	for i := 1; i < numRows; i++ {
		res[i] = make([]int,i+1)
		res[i][0] = 1
		res[i][i] = 1
		for j := 1; j < i; j++ {
			res[i][j] = res[i-1][j-1] + res[i-1][j]
		}
	}

	return res

}

func PrintArray(a [][]int)  {
	for i := 0; i < len(a); i++ {
		fmt.Println(a[i])
	}
}
func main() {
	PrintArray(generate(0))
	PrintArray(generate(1))
	PrintArray(generate(5))
}