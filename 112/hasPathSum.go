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

//给定一个二叉树，它的每个结点都存放一个 0-9 的数字，每条从根到叶子节点的路径都代表一个数字。
//
//例如，从根到叶子节点路径 1->2->3 代表数字 123。
//
//计算从根到叶子节点生成的所有数字之和。
//
//说明: 叶子节点是指没有子节点的节点。
func sumNumbers(root *TreeNode) int {
	sums := make([]int, 0)
	sumNumber(root, &sums)
	sum := 0
	for _, v := range sums {
		sum += v
	}
	return sum
}
func sumNumber(root *TreeNode, sums *[]int) {
	if root == nil {
		return
	}

	if root.Left == nil && root.Right == nil {
		*sums = append(*sums, root.Val)
	} else {
		if root.Left != nil {
			root.Left.Val = root.Left.Val + 10*root.Val
			sumNumber(root.Left, sums)
		}
		if root.Right != nil {
			root.Right.Val = root.Right.Val + 10*root.Val
			sumNumber(root.Right, sums)
		}
	}
	return
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

	val := sumNumbers(root)
	fmt.Println(val)
}
