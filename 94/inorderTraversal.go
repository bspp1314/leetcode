package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func inorderTraversal(root *TreeNode) []int {
	res := make([]int, 0)
	if root == nil {
		return res
	}

	if root.Left != nil {
		left := inorderTraversal(root.Left)
		res = append(res, left...)
	}

	res = append(res, root.Val)

	if root.Right != nil {
		right := inorderTraversal(root.Right)
		res = append(res, right...)
	}
	return res
}
func main() {
	return
}
