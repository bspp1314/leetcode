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

func pathSum(root *TreeNode, targetSum int) [][]int {
	if root == nil {
		return [][]int{}
	}
	res := make([][]int,0)
	var dfs func(node *TreeNode,sub []int,sum int)

	dfs = func(node *TreeNode,sub []int, sum int) {
			if node == nil {
				return
			}

			if node.Right == nil && root.Left == nil {
				if sum == node.Val {
					tem := make([]int,len(sub))
					copy(tem,sub)
					res = append(res,append(tem,root.Val))
				}
				return
			}
			newSub := append(sub,node.Val)
			dfs(node.Left,newSub,sum-node.Val)
			dfs(node.Right,newSub,sum-node.Val)
	}

	dfs(root,[]int{},targetSum)

	return res

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
