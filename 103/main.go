package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func zigzagLevelOrderHelp(root *TreeNode, level int, res *[][]int) {
	if root == nil {
		return
	}

	if len(*res) < level {
		*res = append(*res, []int{root.Val})
	} else {
		if level%2 == 1 {
			(*res)[level-1] = append((*res)[level-1], root.Val)
		} else {
			(*res)[level-1] = append([]int{root.Val}, (*res)[level-1]...)
		}
	}

	zigzagLevelOrderHelp(root.Left, level+1, res)
	zigzagLevelOrderHelp(root.Right, level+1, res)
}

func zigzagLevelOrder(root *TreeNode) [][]int {
	res := make([][]int, 0)
	zigzagLevelOrderHelp(root, 1, &res)
	return res
}

func zigzagLevelOrder2(root *TreeNode) [][]int {
	var res [][]int
	if root == nil {
		return res
	}
	currentLevel := []*TreeNode{root}
	level := 0
	for {
		res = append(res, []int{})
		var nextLevel []*TreeNode
		for i := 0; i < len(currentLevel); i++ {
			node := currentLevel[i]
			if level%2 == 0 {
				res[level] = append(res[level], node.Val)
			} else {
				res[level] = append([]int{node.Val}, res[level]...)
			}
			if node.Left != nil {
				nextLevel = append(nextLevel, node.Left)
			}

			if node.Right != nil {
				nextLevel = append(nextLevel, node.Right)
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
func zigzagLevelOrder3(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}

	nowLevel := []*TreeNode{root}
	res := make([][]int, 0)
	level := 0

	for {
		var nextLevel []*TreeNode
		nowVal := make([]int, 0)

		for i := 0; i < len(nowLevel); i++ {
			nowVal = append(nowVal, nowLevel[i].Val)
			if nowLevel[i].Left != nil {
				nextLevel = append(nextLevel, nowLevel[i].Left)
			}

			if nowLevel[i].Right != nil {
				nextLevel = append(nextLevel, nowLevel[i].Right)
			}
		}

		if level & 1 == 1 {
			left := 0
			right := len(nowVal) -1
			for left <= right {
				nowVal[left],nowVal[right] = nowVal[right],nowVal[left]
				left++
				right--
			}
		}

		level++

		res = append(res, nowVal)

		if len(nextLevel) == 0 {
			break
		} else {
			nowLevel = nextLevel
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

	fmt.Println(zigzagLevelOrder(tree))
	fmt.Println(zigzagLevelOrder2(tree))
	fmt.Println(zigzagLevelOrder2(nil))
}
