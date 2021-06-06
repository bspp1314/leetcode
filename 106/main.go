package main

import (
	"fmt"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func buildTree(inorder []int, postorder []int) *TreeNode {
	tree := &TreeNode{}
	if len(inorder) == 0 || len(postorder) == 0 {
		return nil
	}

	if len(inorder) == 1 {
		tree.Val = inorder[0]
		return tree
	}

	//get root value 
	tree.Val = postorder[len(postorder)-1]

	i := 0
	for i = 0;i<len(inorder);i++ {
		if inorder[i] == tree.Val {
			break
		}
	}


	// get left tree
	tree.Left = buildTree(inorder[0:i],postorder[0:i])
	tree.Right = buildTree(inorder[i+1:],postorder[i:len(postorder)-1])

	return tree


}

func main()  {
		slice := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
		s1 := slice[2:5]
		fmt.Println(s1)
		buildTree([]int{9,3,15,20,7},[]int{9,15,7,20,3})
}


