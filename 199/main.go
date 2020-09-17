package main

import "fmt"

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}


func rightSideView(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}

	currentLevel := []*TreeNode{root}
	res := []int{}

	for  {
		nextLevel := make([]*TreeNode,0)
		res = append(res,currentLevel[len(currentLevel)-1].Val)
		for i:=0;i<len(currentLevel);i++ {
			if currentLevel[i].Left != nil {
				nextLevel = append(nextLevel,currentLevel[i].Left)
			}

			if currentLevel[i].Right != nil {
				nextLevel = append(nextLevel,currentLevel[i].Right)
			}
		}

		currentLevel = nextLevel
		if len(currentLevel) == 0 {
			break
		}
	}

	return res
}

func rightSideView2(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}

	rightNext := rightSideView2(root.Right)
	leftNext := rightSideView2(root.Left)

	res := make([]int,len(rightNext)+1)
	res[0] = root.Val
	for i:= 0;i<len(rightNext);i++ {
		res[i+1] = rightNext[i]
	}


	if len(rightNext) < len(leftNext) {
		for i:= len(rightNext);i < len(leftNext);i++ {
			res = append(res,leftNext[i])
		}
	}

	return res
}

func main() {
	root := &TreeNode{
		Val:   1,
		Left:  &TreeNode{
			Val:   4,
			Left:  &TreeNode{
				Val:   5,
				Left:  &TreeNode{
					Val:   6,
					Left:  &TreeNode{
						Val: 7,
					},
					Right: nil,
				},
				Right: nil,
			},
			Right: nil,
		},
		Right: &TreeNode{
			Val:   2,
			Left:  &TreeNode{
				Val:   3,
				Left:  nil,
				Right: nil,
			},
			Right: nil,
		},
	}

	fmt.Println(rightSideView(root))
	fmt.Println(rightSideView2(root))

}
