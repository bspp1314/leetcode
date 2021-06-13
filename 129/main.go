package main

import (
	"fmt"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

//给定一个二叉树，它的每个结点都存放一个 0-9 的数字，每条从根到叶子节点的路径都代表一个数字。
//
//例如，从根到叶子节点路径 1->2->3 代表数字 123。
//
//计算从根到叶子节点生成的所有数字之和。
//
//说明: 叶子节点是指没有子节点的节点。
func sumNumbers(root *TreeNode) int {
	sum := 0

	var dfs func(subSum int,root *TreeNode)

	dfs = func(subSum int,node *TreeNode)   {
		if node == nil {
			return
		}

		if node.Left == nil && node.Right == nil {
			sum += subSum * 10 + node.Val
		}




		if node.Left != nil {
			dfs(subSum*10+node.Val,node.Left)
		}

		if node.Right != nil {
			dfs(subSum*10+node.Val,node.Right)
		}
	}

	dfs(0,root)
	return sum
}

func main() {
	
	fmt.Println(sumNumbers(&TreeNode{
		Val:   1,
		Left:  &TreeNode{
			Val:   2,
			Left:  &TreeNode{
				Val:   8,
				Left:  nil,
				Right: nil,
			},
			Right: nil,
		},
		Right: &TreeNode{
			Val:   3,
			Left:  nil,
			Right: nil,
		},
	}))
}
