package main


type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func TreeComp(root1,root2 *TreeNode)bool {
	if root1 == nil && root2 == nil {
		return true
	}

	if root1 == nil || root2 == nil {
		return false
	}

	return root1.Val == root2.Val && TreeComp(root1.Left,root2.Right) && TreeComp(root1.Right,root2.Left)
}
func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}
	return TreeComp(root,root)
}

func  isSymmetric2(root *TreeNode)  {

}