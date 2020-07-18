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
	sort.Sort(sort.IntSlice(candidates))
	fmt.Println(candidates)

	dfs(&res, subList, candidates, target, 0)
	return res
}

func dfs(res *[][]int, subList []int, candidates []int, target int, left int) {
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

		if i > left && candidates[i] == candidates[i-1] {
			continue
		}

		dfs(res, append(subList, candidates[i]), candidates, target-candidates[i], i+1)
	}
}

func main() {
	arr := []int{10, 2, 7, 6, 1, 5, 1, 1}
	fmt.Println(arr[2:])
	fmt.Println(combinationSum(arr, 8))

}
