package main

import "fmt"

func permute(nums []int) [][]int {
	if len(nums) == 0 {
		return [][]int{}
	}

	if len(nums) == 1 {
		return [][]int{[]int{nums[0]}}
	}

	res := make([][]int,0)
	for i:=0;i<len(nums);i++ {
		next := make([]int,len(nums)-1)
		copy(next,nums[:i])
		copy(next[i:],nums[i+1:])

		out := permute(nums)
		for _, v := range out {
			res = append(res, append(v, nums[i]))
		}
	}

	return res
}

func permute2(nums []int) [][]int {
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
		next := make([]int,len(nums)-1)
		copy(next,nums[:i])
		copy(next[i:],nums[i+1:])
		dfs(next,res,append(data,nums[i]))

	}

	return
}

func main() {
	fmt.Println(permute2([]int{1,2,3}))
}
