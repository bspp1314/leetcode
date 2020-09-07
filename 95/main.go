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


func inorderTraversal(root *TreeNode) []int {
	// get right
	if root == nil {
		return []int{}
	}
	stack := make([]*TreeNode, 0)
	res := make([]int, 0)
	for root != nil {
		if root.Left != nil  {
			stack = append(stack,root)
			root = root.Left
		}else{
			// stack 5 3 2
			res = append(res,root.Val)
			if root.Right != nil {
				root = root.Right
			}else{
				for {
					if len(stack) == 0 {
						return res
					}else{
						topNode := stack[len(stack)-1]
						if len(stack) == 1 {
							stack = make([]*TreeNode, 0)
						}else{
							stack = stack[:len(stack)-1]
						}
						res = append(res,topNode.Val)
						if topNode.Right != nil {
							root = topNode.Right
							break
						}
					}
				}
			}
		}
	}

	return res
}

func inorderTraversal2(root *TreeNode) []int {
	// get right
	if root == nil {
		return []int{}
	}
	stack := NewStack()
	res := make([]int, 0)
	for root != nil {
		if root.Left != nil  {
			stack.Push(root)
			root = root.Left
		}else{
			// stack 5 3 2
			res = append(res,root.Val)
			if root.Right != nil {
				root = root.Right
			}else{
				for {
					if stack.Len() == 0 {
						return res
					}else{
						topNode := stack.Pop().(*TreeNode)
						res = append(res,topNode.Val)
						if topNode.Right != nil {
							root = topNode.Right
							break
						}
					}
				}
			}
		}
	}

	return res
}
//   1
//
func main() {
	a := []int{1,2,3,4}
	a = a[:len(a)-1]
	fmt.Println(a)
	fmt.Println()
	root := &TreeNode{
		Val: 1,
		Right: &TreeNode{
			Val: 2,
			Right: &TreeNode{
				Val: 4,
				Left: &TreeNode{
					Val: 7,
				},
			},
			Left: &TreeNode{
				Val: 3,
				Right: &TreeNode{
					Val: 6,
				},
				Left: &TreeNode{
					Val: 5,
					Left: &TreeNode{
						Val: 9,
						Left: nil,
						Right: nil,
					},
				},
			},
		},
	}

	out := inorderTraversal(root)
	fmt.Println(out)
	out2 := inorderTraversal2(root)
	fmt.Println(out2)
}
