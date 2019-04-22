package main

import (
	"fmt"
	"log"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

//给定一个二叉树和一个目标和，找到所有从根节点到叶子节点路径总和等于给定目标和的路径。

//说明: 叶子节点是指没有子节点的节点。

//func hasPathSum(root *TreeNode, sum int) bool {
//	if root == nil {
//		return false
//	}
//	if root.Left == nil && root.Right == nil && root.Val == sum {
//		return true
//	} else {
//		return hasPathSum(root.Left, sum-root.Val) || hasPathSum(root.Right, sum-root.Val)
//	}
//}
func pathSum(root *TreeNode, sum int) [][]int {
	pathList := make([]int, 0)
	sums := make([][]int, 0)
	pathSums(root, sum, pathList, &sums)
	return sums
}
func pathSums(root *TreeNode, sum int, pathList []int, sums *[][]int) {
	if root == nil {
		return
	}
	if root.Left == nil && root.Right == nil && root.Val == sum {
		log.Println(sum)
		pathList = append(pathList, root.Val)
		log.Println(pathList)
		*sums = append(*sums, pathList)
	} else {
		pathList = append(pathList, root.Val)
		pathList2 := make([]int, len(pathList))
		copy(pathList2, pathList)
		pathSums(root.Left, sum-root.Val, pathList, sums)
		pathSums(root.Right, sum-root.Val, pathList2, sums)
	}
	return
}

func PrintValue(root *TreeNode) {
	if root == nil {
		return
	}

	fmt.Println(root.Val)
	PrintValue(root.Left)
	PrintValue(root.Right)
}

func main() {
	root := &TreeNode{
		Val: 5,
		Left: &TreeNode{
			Val: 4,
			Left: &TreeNode{
				Val: 11,
				Left: &TreeNode{
					Val:   7,
					Left:  nil,
					Right: nil,
				},
				Right: &TreeNode{
					Val:   2,
					Left:  nil,
					Right: nil,
				},
			},
			Right: nil,
		},
		Right: &TreeNode{
			Val: 8,
			Left: &TreeNode{
				Val:   13,
				Left:  nil,
				Right: nil,
			},
			Right: &TreeNode{
				Val: 4,
				Left: &TreeNode{
					Val:   5,
					Left:  nil,
					Right: nil,
				},
				Right: &TreeNode{
					Val:   1,
					Left:  nil,
					Right: nil,
				},
			},
		},
	}

	val := pathSum(root, 22)
	fmt.Println(val)
}
