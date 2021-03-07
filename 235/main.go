package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func getNext(root,target *TreeNode) (path*TreeNode)  {
		if target.Val < root.Val {
			return root.Left
		} else {
			return root.Right
		}
}


func lowestCommonAncestor(root, p, q *TreeNode) (ans *TreeNode) {

	pathP := root
	pathQ := root

	for pathQ == pathP && pathQ != nil {
		ans = pathQ
		if pathP == p {
			return pathP
		}

		if pathQ == q {
			return pathQ
		}

		pathP  = getNext(pathP,p)
		pathQ =  getNext(pathQ,q)

	}

	return ans
}

func main() {
 tree := &TreeNode{
	 Val:   4,
	 Left:  &TreeNode{
		 Val:   3,
		 Left:  nil,
		 Right: nil,
	 },
	 Right: &TreeNode{
		 Val:   5,
		 Left:  nil,
		 Right: nil,
	 },
 }

 fmt.Println(lowestCommonAncestor(tree,tree,tree.Right))
}
