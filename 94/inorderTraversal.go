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

//左中右
func inorderTraversal2(root *TreeNode) []int {
	res := make([]int, 0)
	if root == nil {
		return res
	}

	var stack []*TreeNode

	for {
		for root != nil {
			stack = append(stack, root)
			root = root.Left
		}

		if len(stack) == 0 {
			break
		}

		c := stack[len(stack)-1]
		res = append(res, c.Val)
		root = root.Right
	}

	return res

}
func main() {
	return
}
