package main

import (
	"container/list"
	"fmt"
	"math"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type Stack struct {
	list *list.List
}

func NewStack() *Stack {
	return &Stack{list.New()}
}
func (s *Stack) Push(v interface{}) {
	s.list.PushBack(v)
}
func (s *Stack) Pop() interface{} {
	e := s.list.Back()
	if e != nil {
		s.list.Remove(e)
		return e.Value
	}
	return nil
}
func (s *Stack) Len() int {
	return s.list.Len()
}

//给定一个二叉树，判断其是否是一个有效的二叉搜索树。
//
// 假设一个二叉搜索树具有如下特征：
//
//
// 节点的左子树只包含小于当前节点的数。
// 节点的右子树只包含大于当前节点的数。
// 所有左子树和右子树自身必须也是二叉搜索树。
//
//
// 示例 1:
//
// 输入:
//    2
//   / \
//  1   3
//输出: true
//
//
// 示例 2:
//
// 输入:
//    5
//   / \
//  1   4
//     / \
//    3   6
//输出: false
//解释: 输入为: [5,1,4,null,null,3,6]。
//     根节点的值为 5 ，但是其右子节点值为 4 。
func isValidBST(root *TreeNode) bool {
	return isValidBSTHelp(root, math.MinInt64 , math.MaxInt64)
}

func isValidBSTHelp(root *TreeNode, min, max int)bool{
	if root == nil{
		return true
	}
	if root.Val <= min || root.Val >= max{
		return false
	}

	return isValidBSTHelp(root.Left, min, root.Val) && isValidBSTHelp(root.Right, root.Val, max)
}
func main() {
	tree := &TreeNode{
		Val:   10,
		Left:  &TreeNode{
			Val:   10,
			Left:  nil,
			Right: nil,
		},
		Right: &TreeNode{
			Val:   15,
			Left:  &TreeNode{
				Val:   5,
				Left:  nil,
				Right: nil,
			},
			Right: &TreeNode{
				Val:   20,
				Left:  nil,
				Right: nil,
			},
		},
	}

	fmt.Println(isValidBST(tree))
}
