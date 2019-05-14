package main

import "fmt"

/**
* Definition for a binary tree node.*/
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func allPossibleFBT(N int) []*TreeNode {
	res := make([]*TreeNode, 0)
	if N%2 == 0 {
		return nil
	} else if N == 1 {
		root := &TreeNode{Val: 0}
		res = append(res, root)
	} else {
		for i := 1; i < N/2+1; i += 2 {
			left := allPossibleFBT(i)
			right := allPossibleFBT(N - i - 1)
			for _, k := range left {
				for _, m := range right {
					root := &TreeNode{Val: 0}
					root.Left = k
					root.Right = m
					res = append(res, root)
					if i != (N - i - 1) {
						root2 := &TreeNode{Val: 0}
						root2.Left = m
						root2.Right = k
						res = append(res, root2)

					}
				}
			}
		}
	}
	return res
}
func main() {
	allPossibleFBT(19)
	fmt.Println("vim-go")
}
