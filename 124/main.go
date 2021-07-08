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


func maxPathSum(root *TreeNode) int {
	maxSum := math.MinInt64
	var subMaxPahSum func(*TreeNode) int
	max := func(a, b int) int {
		if a > b {
			return a
		}

		return b
	}

	subMaxPahSum = func(node *TreeNode) int {
		if node == nil {
			return 0
		}

		leftVal := max(subMaxPahSum(node.Left), 0)
		rightVal := max(subMaxPahSum(node.Right), 0)

		newMaxSum := node.Val + leftVal + rightVal
		maxSum = max(newMaxSum, maxSum)

		return node.Val + max(leftVal, rightVal)
	}

	subMaxPahSum(root)
	return maxSum

}



func main() {
	root := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: -2,
			Left: &TreeNode{
				Val: 1,
				Left: &TreeNode{
					Val:   -2,
					Left:  nil,
					Right: nil,
				},
				Right: nil,
			},
			Right: &TreeNode{
				Val:   3,
				Left:  nil,
				Right: nil,
			},
		},
		Right: &TreeNode{
			Val: -3,
			Left: &TreeNode{
				Val:   -2,
				Left:  nil,
				Right: nil,
			},
			Right: nil,
		},
	}

	out := maxPathSum(root)
	fmt.Println(out)
	//left
	//right
	fmt.Println(maxPathSum(root.Left))
}
