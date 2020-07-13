package main

import (
	"fmt"
	"sort"
)

func combinationSum(candidates []int, target int) [][]int {
	l := len(candidates)

	res := make([][]int, 0)
	if l == 0 {
		return res
	}

	subList := make([]int, 0)
	fmt.Println(candidates)
	sort.Sort(sort.IntSlice(candidates))

	combinationSumHelp(&res, subList, candidates, target, 0)
	return res
}

func combinationSumHelp(res *[][]int, subList []int, candidates []int, target int, left int) {
	if target == 0 {
		tmp := make([]int, len(subList))
		copy(tmp, subList)
		*res = append(*res, tmp)
		return
	}

	for i := left; i < len(candidates); i++ {
		if candidates[i] > target { // 剪枝操作，当当前数值大于目标值，则后续无需遍历
			return
		}

		combinationSumHelp(res, append(subList, candidates[i]), candidates, target-candidates[i], i)
	}
}

func main() {
	arr := []int{2, 3, 5}
	fmt.Println(combinationSum(arr, 8))
}
