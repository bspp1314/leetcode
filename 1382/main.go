package  main

import (
	"fmt"
)

type TreeNode struct {
	Left *TreeNode
	Right *TreeNode
	Val int
}


func balanceBST(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}

	res := make([]*TreeNode,0)
	stack := make([]*TreeNode,0)
	current := root

	for current != nil ||  len(stack) != 0 {
		for current != nil {
			stack = append(stack,current)
			current = current.Left
		}
		current = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		res = append(res, current)
		current = current.Right
	}

	return balanceBSTHelp(res)
}



func balanceBSTHelp(nodes []*TreeNode)*TreeNode {
	if len(nodes) == 0 {
		return nil
	}

	if len(nodes) == 1 {
		nodes[0].Left = nil
		nodes[0].Right = nil
		return nodes[0]
	}

	mid := len(nodes) / 2
	root := nodes[mid]

	root.Left = balanceBSTHelp(nodes[:mid])

	if (mid + 1) < len(nodes) {
		root.Right = balanceBSTHelp(nodes[mid+1:])
	}else{
		root.Right = nil
	}

	return root
}

func inorderTraversal(root *TreeNode) []int {
	res := make([]int, 0)
	stack := make([]*TreeNode,0)
	current := root
	for current != nil || len(stack) != 0 {
		for current != nil {
			stack = append(stack,current)
			current = current.Left
		}
		current = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		res = append(res, current.Val)
		current = current.Right
	}
	return res
}

func main() {
	tree := &TreeNode{
		Left:  nil,
		Right: &TreeNode{
			Left:  nil,
			Right: &TreeNode{
				Left:  nil,
				Right: &TreeNode{
					Left:  nil,
					Right: nil,
					Val:   4,
				},
				Val:   3,
			},
			Val:   2,
		},
		Val:   1,
	}
	out := balanceBST(tree)
	fmt.Println("out",out)
	fmt.Println("out left",out.Left)
	fmt.Println("out left right",out.Left.Right)
	fmt.Println("out left left ",out.Left.Left)
	fmt.Println(out.Right)
	fmt.Println(inorderTraversal(out))
}
