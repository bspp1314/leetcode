package main

import (
	"container/list"
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

func preorderTraversal(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}

	res := make([]int, 0)
	stack := NewStack()
	current := root

	stack.Push(current)
	for stack.Len() != 0 {
		current = stack.Pop().(*TreeNode)
		res = append(res, current.Val)
		if current.Right != nil {
			stack.Push(current.Right)
		}
		if current.Left != nil {
			stack.Push(current.Left)
		}
	}
	return res

}

//递归法
func preorderTraversal2(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}

	left := preorderTraversal2(root.Left)
	right := preorderTraversal2(root.Right)
	left = append(left, right...)
	left = append(left, root.Val)
	return left
}
func main() {
	return
}
