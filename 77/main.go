package main

import "fmt"

//[
//[3],
//[1],
//[2],
//[1,2,3],
//[1,3],
//[2,3],
//[1,2],
//[]
//]

/// 0 1
//  1 2
//  2 4
//  3 8
func subsets(nums []int) [][]int {
	if len(nums) == 0 {
		return [][]int{[]int{}}
	}

	subRes := subsets(nums[1:])
	res := make([][]int,1 << len(nums))

	j:= 0
	for i:= 0;i<len(subRes);i++ {
		res[j] = subRes[i]
		dest := make([]int,len(subRes[i]))
		copy(dest,subRes[i])
		res[j+1] = append(dest,nums[0])
		j = j + 2
	}


	return res
}

func main() {
	res := subsets([]int{1,2,3,4,5})
	fmt.Println(res)
}
