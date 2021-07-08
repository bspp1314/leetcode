package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
func preorderTraversal(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}



	stack := []*TreeNode{root}
	res := make([]int,0)

	for len(stack) != 0 {
		root = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		res = append(res,root.Val)
		if root.Right != nil {
			stack = append(stack,root.Right)
		}

		if root.Left != nil {
			stack = append(stack,root.Left)
		}
	}

	return res
}

func main() {
	return
}
