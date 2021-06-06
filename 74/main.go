package main

import (
	"fmt"
)


func searchMatrix(matrix [][]int, target int) bool {
	if len(matrix) == 0 || len(matrix[0]) == 0  {
		return false
	}



	col := len(matrix)
	row := len(matrix[0])

	i := col -1
	j := 0

	for i >= 0 && j < row {
		if matrix[i][j] > target {
			i--
		}else if matrix[i][j] < target{
			j++
		}else{
			return true
		}
	}

	return false
}

func main() {
  arrays := [][]int{
	  []int{1,2,3,4},
	  []int{6,8,10,12},
	  []int{13,14,19,20},
  }

	for i:= 1;i <= 10;i++ {
		fmt.Println("i is ",i,"in",searchMatrix(arrays,i))
	}
}
