package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isBalanced(root *TreeNode) bool {
	fmt.Println(getRootHeight(root))
	return getRootHeight(root) != -1
}

func getRootHeight(root *TreeNode) int  {
	if root == nil {
		return 0
	}

	// get left
	leftH := getRootHeight(root.Left)
	if leftH == -1 {
		return -1
	}

	rightH := getRootHeight(root.Right)
	if rightH == -1 {
		return -1
	}

	if rightH > leftH {
		v := rightH - leftH
		if v > 1 {
			return -1
		}
		return rightH+1
	}

	v := leftH - rightH
	if v > 1 {
		return -1
	}

	return leftH+1
}

func main() {
	tree := &TreeNode{
		Val:1,
		Left: &TreeNode{
			Val:   2,
			Left:  &TreeNode{
				Val:   3,
				Left:  &TreeNode{
					Val:   4,
					Left:  nil,
					Right: nil,
				},
				Right: &TreeNode{
					Val:   4,
					Left:  nil,
					Right: nil,
				},
			},
			Right: &TreeNode{
				Val: 3,
			},
		},
		Right: &TreeNode{
			Val:   2,
			Left:  nil,
			Right: nil,
		},
	}

	isBalanced(tree)
}