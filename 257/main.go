package main

import (
	"fmt"
	"strings"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func binaryTreePaths(root *TreeNode) (res []string) {
	if root == nil {
		return []string{}
	}

	var dfs func(parent []string, node *TreeNode)

	dfs = func(parent []string, node *TreeNode) {
		if node == nil {
			return
		}

		if node.Right == nil && node.Left == nil {
			parent = append(parent, fmt.Sprintf("%d", node.Val))
			res = append(res, strings.Join(parent, "->"))
			return
		}

		if node.Left != nil {
			dfs(append(parent, fmt.Sprintf("%d", node.Val)), node.Left)
		}

		if node.Right != nil {
			dfs(append(parent, fmt.Sprintf("%d", node.Val)), node.Right)
		}
	}

	dfs([]string{}, root)

	return res

}

func main() {
	n := &TreeNode{
		Val: 2,
		Left: &TreeNode{
			Val:  0,
			Left: nil,
			Right: &TreeNode{
				Val:   1,
				Left:  nil,
				Right: nil,
			},
		},
		Right: &TreeNode{
			Val:   4,
			Left:  nil,
			Right: nil,
		},
	}

	fmt.Println(binaryTreePaths(n))
}
