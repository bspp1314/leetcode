package main

import (
	"fmt"
	"sort"
)

func combinationSum3(k int, n int) [][]int {
	res := make([][]int,0)
	traceBack(k,1,n,&res,[]int{},0)
	return res
}

func traceBack(k int,left int,n int,res *[][]int,temp []int,sum int){
	if k == 0 {
		if sum == n {
			temp2 := make([]int,len(temp))
			copy(temp2,temp)
			*res = append(*res,temp2)
		}
		return
	}


	for i := left; i < n;i++ {
		if len(temp) != 0 {
			if temp[len(temp)-1] >= i {
				continue
			}
		}

		traceBack(k-1,i,n,res,append(temp,i),sum+i)
	}
}

func combinationSum2(candidates []int, target int) [][]int {
	sort.Sort(sort.IntSlice(candidates))
	res := make([][]int,0)
	Backtrack(0,candidates,make([]int,0),target,&res)
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

func Backtrack(res *[][]int,left int,candidates []int,tem []int,target int) {
	if target == 0 {
		tem2 := make([]int,len(tem))
		copy(tem2,tem)
		*res = append(*res,tem2)
	}

	for i:=left;i < len(candidates);i++ {
		if candidates[i] > target {
			return
		}

		if i > left && candidates[i] == candidates[i-1] {
			continue
		}

		Backtrack(res,left+1,candidates,append(tem,candidates[i]),target-candidates[i])
	}
}

func main() {
	out := combinationSum3(4,20)
	fmt.Println(out)
}