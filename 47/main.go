package main

import (
	"fmt"
	"sort"
)

func permuteUnique(nums []int) [][]int {
	sort.Sort(sort.IntSlice(nums))
	res := make([][]int,0)
	dfs(nums,&res,[]int{})
	return res
}

func dfs(nums []int,res *[][]int,data []int){
	if len(nums) == 1 {
		tem := make([]int,len(data)+1)
		copy(tem,data)
		tem[len(tem)-1] = nums[0]
		*res = append(*res,tem)
		return
	}

	for i:=0;i<len(nums);i++ {
		if i != 0 {
			if nums[i] == nums[i-1] {
				continue
			}
		}
		next := make([]int,len(nums)-1)
		copy(next,nums[:i])
		copy(next[i:],nums[i+1:])
		dfs(next,res,append(data,nums[i]))
	}

	return
}

func main() {
	out := permuteUnique([]int{1,1,2})
	fmt.Println(out)
}
