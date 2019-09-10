package main

import (
	"container/list"
	"fmt"
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
	res := make([]int, 0)
	stack := NewStack()
	s := list.New()
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
func postorderTraversla(root *TreeNode) []int {
	res := make([]int, 0)
	stack := NewStack()
	current := root

	for stack.Len() != 0 {
		stack.Push(current)
		res = append([]int{current.Val}, res...)
		if current.Right != nil {
			stack.Push(current.Right)
		}
		if current.Left != nil {
			stack.Push(current.Left)
		}

	}
	return res
}
func main() {
	root := &TreeNode{
		Val: 10,
		Left: &TreeNode{
			Val:   2,
			Left:  nil,
			Right: nil,
		},
		Right: &TreeNode{
			Val: 3,
		},
	}
	fmt.Println(preorderTraversal(root))
	return
}
