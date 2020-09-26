package main

import (
	"fmt"
	"sort"
)

func subsetsWithDup(nums []int) [][]int {
	sort.Sort(sort.IntSlice(nums))
	res := make([][]int,0)
	dfs(nums,0,[]int{},&res)
	return res
}

func subsets(nums []int) [][]int {
	if len(nums) == 0 {
		return [][]int{[]int{}}
	}

	s := make([][]int, 0)

	i := 0
	for i = 0; i < len(nums); i++ {
		if i != 0 {
			if nums[i] != nums[i-1] {
				break
			}else{
				tem := make([]int, len(nums[0:i+1]))
				copy(tem, nums[0:i+1])
				s = append(s, tem)
			}
		} else {
			tem := make([]int, len(nums[0:i+1]))
			copy(tem, nums[0:i+1])
			s = append(s, tem)
		}
	}


	if i == len(nums) {
		s = append(s, []int{})
		return s
	} else {
		subRes := subsets(nums[i:])
		res := make([][]int, 0, len(s))
		for j := 0; j < len(subRes); j++ {
			res = append(res, subRes[j])

			for k := 0; k < len(s); k++ {
				tem := make([]int, len(subRes[j])+len(s[k]))
				copy(tem, subRes[j])
				copy(tem[len(subRes[j]):], s[k])
				res = append(res, tem)
			}
		}

		return res
	}
}

func dfs(nums []int,start int,temp []int,res *[][]int)  {
	*res = append(*res,temp)
	for i:= start;i < len(nums);i++ {
	 	if i > start && nums[i] == nums[i-1] {
	 		continue
		}
		newTemp := make([]int,len(temp))
		copy(newTemp,temp)
		dfs(nums,i+1,append(newTemp,nums[i]),res)
	 }
}
func generateParenthesis(n int) []string {
	res := make([]string,0)
	dfsG(0,0,n,"",&res)
	return res
}

func dfsG(leftNum,rightNum,n int,tem string,res *[]string){
	if leftNum == n && rightNum == n {
		*res = append(*res,tem)
		return
	}

	if leftNum == n {
		dfsG(leftNum,rightNum+1,n,tem+")",res)
		return
	}

	if leftNum > rightNum {
		dfsG(leftNum+1,rightNum,n,tem+"(",res)
		dfsG(leftNum,rightNum+1,n,tem+")",res)
		return
	}

	if leftNum == rightNum {
		dfsG(leftNum+1,rightNum,n,tem+"(",res)
		return
	}
}

func main() {
	out := subsetsWithDup([]int{1, 1, 2,2})
	fmt.Println(out)
	outin := generateParenthesis(1)
	fmt.Println(outin)
}
