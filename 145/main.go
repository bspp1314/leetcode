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

	stack := make([]*TreeNode,0)
	//用来记录左节点是否遍历过
	var pre *TreeNode

	for root !=nil || len(stack) != 0 {
		for root != nil {
			stack = append(stack,root)
			root = root.Left
		}

		n := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		if n.Right == nil || n.Right == pre {
			pre = n
			root = nil
			res = append(res,n.Val)
		}else{
			stack = append(stack,n)
			root = n.Right
		}
	}

	return res
}
func main()  {

}
