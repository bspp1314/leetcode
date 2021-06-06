package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func pathSum(root *TreeNode, targetSum int) [][]int {
	if root == nil {
		return [][]int{}
	}
	res := make([][]int,0)
	var dfs func(node *TreeNode,sub []int,sum int)

	dfs = func(node *TreeNode,sub []int, sum int) {
		if node == nil {
			return
		}


		if node.Right == nil && node.Left == nil {
			if sum == node.Val {
				tem := make([]int,len(sub))
				copy(tem,sub)
				res = append(res,append(tem,node.Val))
			}
			return
		}
		newSub := append(sub,node.Val)
		dfs(node.Left,newSub,sum-node.Val)
		dfs(node.Right,newSub,sum-node.Val)
	}

	dfs(root,[]int{},targetSum)

	return res

}

func main() {
}
