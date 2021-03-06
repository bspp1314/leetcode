package main

import (
	"fmt"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}


func levelOrder(root *TreeNode) [][]int {
	res := make([][]int, 0)
	levelOrderHelp(root, 1, &res)
	return res
}

func levelOrderHelp(root *TreeNode, level int, res *[][]int) {
	if root == nil {
		return
	}

	if len(*res) < level {
		*res = append(*res, []int{root.Val})
	} else {
		(*res)[level-1] = append((*res)[level-1], root.Val)
	}

	levelOrderHelp(root.Left, level+1, res)
	levelOrderHelp(root.Right, level+1, res)
}

func levelOrder2(root *TreeNode)[][]int  {
	var res [][]int
	if root == nil {
		return res
	}
	currentLevel := []*TreeNode{root}
	level := 0
	for {
		res = append(res,[]int{})
		var nextLevel []*TreeNode
		for i:=0;i<len(currentLevel);i++ {
			node := currentLevel[i]
			res[level] = append(res[level],node.Val)

			if node.Left != nil {
				nextLevel = append(nextLevel,node.Left)
			}

			if node.Right != nil {
				nextLevel = append(nextLevel,node.Right)
			}
		}

		level++
		currentLevel = nextLevel
		if len(currentLevel) == 0 {
			break
		}
	}

	return res
}

func main() {
	tree := &TreeNode{
		Val: 10,
		Left: &TreeNode{
			Val: 11,
			Left: &TreeNode{
				Val:   12,
				Left:  nil,
				Right: nil,
			},
			Right: nil,
		},
		Right: &TreeNode{
			Val: 20,
			Left: &TreeNode{
				Val:   21,
				Left:  nil,
				Right: nil,
			},
			Right: nil,
		},
	}
	fmt.Println(levelOrder(tree))
	return
}
