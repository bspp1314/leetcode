package main

import (
	"fmt"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

//给定一个二叉树和一个目标和，判断该树中是否存在根节点到叶子节点的路径，这条路径上所有节点值相加等于目标和。
//说明: 叶子节点是指没有子节点的节点。
func hasPathSum(root *TreeNode, sum int) bool {
	if root == nil {
		return false
	}
	if root.Left == nil && root.Right == nil && root.Val == sum {
		return true
	} else {
		return hasPathSum(root.Left, sum-root.Val) || hasPathSum(root.Right, sum-root.Val)
	}
}

func PrintValue(root *TreeNode) {
	if root == nil {
		return
	}

	fmt.Println(root.Val)
	PrintValue(root.Left)
	PrintValue(root.Right)
}

func main() {
	root := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 2,
			//Left: &TreeNode{
			//	Val: 1,
			//	Left: &TreeNode{
			//		Val:   4,
			//		Left:  nil,
			//		Right: nil,
			//	},
			//	Right: &TreeNode{
			//		Val:   5,
			//		Left:  nil,
			//		Right: nil,
			//	},
			//},
			Right: nil,
		},
		Right: &TreeNode{
			Val:   3,
			Left:  nil,
			Right: nil,
		},
	}

	val := hasPathSum(root, 4)
	fmt.Println(val)
}
