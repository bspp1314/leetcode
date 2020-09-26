package main

type TreeNode struct {
	Left *TreeNode
	Right *TreeNode
	Val int
}

func findMode(root *TreeNode) (answer []int) {
	base := 0
	count := 0
	maxCount := 0
	res := make([]int,0)
	dfs(root,&base,&count,&maxCount,&res)

	return res
}

func dfs(root *TreeNode,base *int,count *int,maxCount *int,res *[]int) {
	if root == nil {
		return
	}

	dfs(root.Left,base,count,maxCount,res)
	if root.Val == *base {
		*count++
	}else {
		*base = root.Val
		*count = 1
	}
	if *count == *maxCount {
		*res = append(*res,*base)
	}else if *count > *maxCount {
		*res = []int{*base}
		*maxCount = *count
	}
	dfs(root.Right,base,count,maxCount,res)
}

func main() {
}