package main

import (
	"fmt"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func flatten(root *TreeNode) {
	if root == nil {
		return
	}

	if root.Left == nil && root.Right == nil {
		return
	} else {
		left := root.Left
		right := root.Right
		flatten(left)
		flatten(right)

		root.Left = nil
		if left == nil {
			root.Right = right

		} else if right == nil {
			root.Right = left
		} else {
			root.Right = left
			last := left
			for {
				if last.Right == nil {
					break
				} else {
					last = last.Right
				}
			}
			last.Right = right
		}
	}

}
func flatten2(root *TreeNode) {
	if root == nil {
		return
	}

	if root.Left == nil && root.Right == nil {
		return
	} else {
		flatten(root.Left)
		flatten(root.Right)
		if root.Left != nil {
			right := root.Right
			root.Right = root.Left
			root.Left = nil

			last := root.Right
			for last.Right != nil {
				last = last.Right
			}
			last.Right = right
		}
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
		Val: 5,
		Left: &TreeNode{
			Val: 4,
			Left: &TreeNode{
				Val: 11,
				Left: &TreeNode{
					Val:   7,
					Left:  nil,
					Right: nil,
				},
				Right: &TreeNode{
					Val:   2,
					Left:  nil,
					Right: nil,
				},
			},
			Right: nil,
		},
		Right: &TreeNode{
			Val: 8,
			Left: &TreeNode{
				Val:   13,
				Left:  nil,
				Right: nil,
			},
			Right: &TreeNode{
				Val: 4,
				Left: &TreeNode{
					Val:   5,
					Left:  nil,
					Right: nil,
				},
				Right: &TreeNode{
					Val:   1,
					Left:  nil,
					Right: nil,
				},
			},
		},
	}
	PrintValue(root)
	fmt.Printf("||||||||||||||||||||||||||||||||||||||||||\n")
	flatten(root)
	PrintValue(root)
}
