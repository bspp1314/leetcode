package main



func getIndex(v int,col int) (i,j int ) {
	i = v /  col
	j = v %  col
	return i,j
}

func searchMatrix(matrix [][]int, target int) bool {
	if len(matrix) == 0 || len(matrix[0]) == 0  {
		return false
	}



	row := len(matrix)
	col := len(matrix[0])
	left := 0
	right := row * col  - 1



	for left <= right {
		mid := left + ((right - left)>>1)
		i,j := getIndex(mid,col)
		if matrix[i][j] > target {
			right = mid - 1
		}else if matrix[i][j] < target {
			left = mid + 1
		}else{
			return true
		}
	}

	return false
}

func main() {

}
