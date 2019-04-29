package main

import (
	"fmt"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isCompleteTree(root *TreeNode) bool {
	if root == nil {
		return true
	}

	if root.Left == nil && root.Right == nil {
		return true
	}
	return nil

}
func help(root *TreeNode) (bool, int) {
	if root == nil {
		return true, 0
	}

	if root.Left == nil && root.Right == nil {
		return true, 1
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
	//root := &TreeNode{
	//	Val: 5,
	//	Left: &TreeNode{
	//		Val: 4,
	//		Left: &TreeNode{
	//			Val: 11,
	//			Left: &TreeNode{
	//				Val:   7,
	//				Left:  nil,
	//				Right: nil,
	//			},
	//			Right: &TreeNode{
	//				Val:   2,
	//				Left:  nil,
	//				Right: nil,
	//			},
	//		},
	//		Right: nil,
	//	},
	//	Right: &TreeNode{
	//		Val: 8,
	//		Left: &TreeNode{
	//			Val:   13,
	//			Left:  nil,
	//			Right: nil,
	//		},
	//		Right: &TreeNode{
	//			Val: 4,
	//			Left: &TreeNode{
	//				Val:   5,
	//				Left:  nil,
	//				Right: nil,
	//			},
	//			Right: &TreeNode{
	//				Val:   1,
	//				Left:  nil,
	//				Right: nil,
	//			},
	//		},
	//	},
	//}
	//PrintValue(root)
	//fmt.Printf("||||||||||||||||||||||||||||||||||||||||||\n")
	//flatten(root)
	//PrintValue(root)
}
