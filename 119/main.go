package main

import "fmt"

func getRow(rowIndex int) []int {
	if rowIndex == 0 {
		return []int{1}
	}

	if rowIndex == 1 {
		return []int{1, 1}
	}

	resPre := make([]int, rowIndex)
	res := make([]int, rowIndex+1)
	resPre[0] = 1
	res[0] = 1
	res[1] = 1

	for i := 2; i < rowIndex + 1; i++ {
		copy(resPre,res[:i])
		res[0] = 1
		res[i] = 1

		for j := 1; j < i; j++ {
			res[j] = resPre[j-1] + resPre[j]
		}
	}

	return res

}
func main() {
	fmt.Println(getRow(3))
}
