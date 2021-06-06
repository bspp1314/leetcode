package main

type TreeNode struct {
	Left  *TreeNode
	Right *TreeNode
	Val   int
}


//  1 . 从根节点开始

func morrisInOrderTraversal(root *TreeNode) []int {
	node := root
	var prev *TreeNode

	res := make([]int,0)
	for node != nil {
		if node.Left == nil {
			res = append(res,node.Val)
			node = node.Right
		}else{
			//查询前驱节点
			prev = node.Left

		}
	}

	return res
}

func main() {

}
