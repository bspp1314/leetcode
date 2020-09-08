package main

import "fmt"

//二叉树的中序遍历
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
type Res struct{
	Values [][]int
}
func levelOrder(root *TreeNode) [][]int {
	res :=  make([][]int,0)
	levelOrderHelp(root,1,&res)
	return res
}


func levelOrderHelp(root *TreeNode,level int,res *[][]int){
	if root == nil {
		return
	}

	if len(*res) < level {
		*res = append(*res,[]int{root.Val})
	}else{
		(*res)[level-1] = append((*res)[level-1],root.Val)
	}

	levelOrderHelp(root.Left,level+1,res)
	levelOrderHelp(root.Right,level+1,res)
}
func main() {
	tree := &TreeNode{
		Val:   10,
		Left:  &TreeNode{
			Val:   11,
			Left:  &TreeNode{
				Val:   12,
				Left:  nil,
				Right: nil,
			},
			Right: nil,
		},
		Right: &TreeNode{
			Val:   20,
			Left:  &TreeNode{
				Val:   21,
				Left:  nil,
				Right: nil,
			},
			Right: nil,
		},
	}
	fmt.Println(levelOrder(tree))
	return
}
