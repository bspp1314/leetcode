package main

import "fmt"

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

// 广度优先遍历
func rightSideView(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}

	currentLevel := []*TreeNode{root}
	var res []int

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


// 深度优先遍历
func rightSideView2(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}

	if root.Left == nil && root.Right == nil {
		return []int{root.Val}
	}

	right := rightSideView(root.Right)
	left  := rightSideView(root.Left)

	res := []int{root.Val}
	res = append(res,right...)
	if len(right) < len(left) {
		res = append(res,left[len(right):]...)
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
