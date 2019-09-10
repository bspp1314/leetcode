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

func inorderTraversal2(root *TreeNode) []int {
	res := make([]int, 0)
	stack := NewStack()
	current := root
	for current != nil || stack.Len() != 0 {
		for current != nil {
			stack.Push(current)
			current = current.Left
		}
		current = stack.Pop().(*TreeNode)
		res = append(res, current.Val)
		current = current.Right
	}
	return res

}
func inorderTraversal(root *TreeNode) []int {
	res := make([]int, 0)
	if root == nil {
		return res
	}

	if root.Left != nil {
		left := inorderTraversal(root.Left)
		res = append(res, left...)
	}

	res = append(res, root.Val)

	if root.Right != nil {
		right := inorderTraversal(root.Right)
		res = append(res, right...)
	}
	return res
}
func main() {
	s := NewStack()
	s.Push(10)
	v1 := s.Pop().(int)
	fmt.Println(v1)
	//root := &TreeNode{
	//	Val: 10,
	//	Left: &TreeNode{
	//		Val:   2,
	//		Left:  nil,
	//		Right: nil,
	//	},
	//	Right: &TreeNode{
	//		Val: 3,
	//	},
	//}
	//fmt.Println(inorderTraversal2(root))
	return
}
