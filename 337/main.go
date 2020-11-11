package main

import (
	"fmt"
)

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func Max(a,b int) int  {
	if a > b {
		return a
	}

	return b
}
func rob(root *TreeNode) int {
	v := robHelp(root)
	return Max(v[0],v[1])
}

func robHelp(root *TreeNode) []int  {
	if root == nil {
		return []int{0,0}
	}
	left := robHelp(root.Left)
	right := robHelp(root.Right)

	// steal
	res1 := root.Val + left[1] + right[1]
	// no steal
	res2 := Max(left[0],left[1]) + Max(left[0],left[1])

	return []int{res1,res2}
}

func main() {
	r := &TreeNode{
		Val:   3,
		Left:  &TreeNode{
			Val:   2,
			Left:  nil,
			Right: &TreeNode{
				Val:   3,
				Left:  nil,
				Right: nil,
			},
		},
		Right: &TreeNode{
			Val:   3,
			Left:  nil,
			Right: &TreeNode{
				Val:   1,
				Left:  nil,
				Right: nil,
			},
		},
	}
	fmt.Println(rob(r))
}
