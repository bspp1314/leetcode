package main

//前 中左右
//中 左中右
//后 左右中

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func postorderTraversal(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}

	var res []int

	stack := make([]*TreeNode, 0)
	//用来记录上一个遍历节点的位置
	var pre *TreeNode

	for root != nil || len(stack) != 0 {
		for root != nil {
			stack = append(stack, root)
			root = root.Left
		}

		root = stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		// 判断Right 是否被遍历过
		if root.Right == nil || root.Right == pre {
			res = append(res, root.Val)
			pre = root
			root = nil
		} else {
			// root.Right 没有被遍历过
			stack = append(stack, root)
			root = root.Right
		}

	}

	return res
}
func main() {

}
