package main

import "fmt"

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0

	}

	leftMaxDepth := maxDepth(root.Left) + 1
	rightMaxDepth := maxDepth(root.Right) + 1

	if leftMaxDepth > rightMaxDepth {
		return leftMaxDepth

	}

	return rightMaxDepth

}

func maxDepth2(root *TreeNode) int {
	if root == nil {
		return 0

	}
	queue := []*TreeNode{}
	queue = append(queue, root)
	ans := 0
	for len(queue) > 0 {
		sz := len(queue)
		for sz > 0 {
			node := queue[0]
			queue = queue[1:]
			if node.Left != nil {
				queue = append(queue, node.Left)

			}
			if node.Right != nil {
				queue = append(queue, node.Right)

			}
			sz--

		}
		ans++

	}
	return ans

}

func main() {
	fmt.Println("vim-go")
}
