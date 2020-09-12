package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSameTree(p *TreeNode, q *TreeNode) bool {
	if p == nil {
		return p == q
	}

	if q == nil {
		return p == q
	}

	return p.Val == q.Val && isSameTree(p.Left,q.Left) && isSameTree(p.Right,q.Right)
}

func main() {

}
