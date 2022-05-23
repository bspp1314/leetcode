package main

import (
	"fmt"
	"math"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isValidBST(root *TreeNode) bool {

	return isValidBSTHelp(root, math.MinInt64, math.MaxInt64)
}

// 如果该二叉树的左子树不为空，则左子树上所有节点的值均小于它的根节点的值；
//若它的右子树不空，则右子树上所有节点的值均大于它的根节点的值；它的左右子树也为二叉搜索树。
func isValidBSTHelp(root *TreeNode, min, max int) bool {
	if root == nil {
		return true
	}
	if root.Val <= min || root.Val >= max {
		return false
	}

	return isValidBSTHelp(root.Left, min, root.Val) && isValidBSTHelp(root.Right, root.Val, max)
}

func main() {
	tree := &TreeNode{
		Val: 10,
		Left: &TreeNode{
			Val:   10,
			Left:  nil,
			Right: nil,
		},
		Right: &TreeNode{
			Val: 15,
			Left: &TreeNode{
				Val:   5,
				Left:  nil,
				Right: nil,
			},
			Right: &TreeNode{
				Val:   20,
				Left:  nil,
				Right: nil,
			},
		},
	}

	fmt.Println(isValidBST(tree))
}
