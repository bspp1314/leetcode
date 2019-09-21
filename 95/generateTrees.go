package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func NewTreeNode(val int) *TreeNode {
	return &TreeNode{
		Val:   val,
		Left:  nil,
		Right: nil,
	}
}

func generateTrees(n int) []*TreeNode {
	dp := make([][]*TreeNode, n+1)
	return generateTreesHelp(1, n, dp)
}

func generateTreesHelp(start, end int, dp [][]*TreeNode) []*TreeNode {
	if start > end {
		return nil
	}

	if dp[end-start+1] != nil {
		return Clones(dp[end-start+1], start-dp[end-start+1][0].Val)
	}

	res := make([]*TreeNode, 0)
	for i := start; i <= end; i++ {
		left := generateTreesHelp(start, i-1, dp)
		right := generateTreesHelp(i+1, end, dp)
		if len(left) == 0 && len(right) == 0 {
			node := NewTreeNode(i)
			res = append(res, node)
		} else if len(left) == 0 {
			for j := 0; j < len(right); j++ {
				node := NewTreeNode(i)
				node.Right = right[j]
				res = append(res, node)
			}
		} else if len(right) == 0 {
			for k := 0; k < len(left); k++ {
				node := NewTreeNode(i)
				node.Left = left[k]
				res = append(res, node)
			}
		} else {
			for j := 0; j < len(right); j++ {
				for k := 0; k < len(left); k++ {
					node := NewTreeNode(i)
					node.Right = right[j]
					node.Left = left[k]
					res = append(res, node)
				}
			}

		}
	}
	dp[end-start+1] = res
	return res
}

func Clones(roots []*TreeNode, offset int) []*TreeNode {
	if offset == 0 {
		return roots
	}
	res := make([]*TreeNode, 0)
	for i := 0; i < len(roots); i++ {
		res = append(res, Clone(roots[i], offset))
	}
	return res
}

func Clone(root *TreeNode, offset int) *TreeNode {
	if root == nil {
		return nil
	}
	newRoot := NewTreeNode(root.Val + offset)
	newRoot.Left = Clone(root.Left, offset)
	newRoot.Right = Clone(root.Right, offset)
	return newRoot
}

func main() {
	get := generateTrees(5)
	fmt.Println(get)
	fmt.Println("vim-go")
}
