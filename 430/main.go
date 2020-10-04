package  main

import "fmt"

type Node struct {
     Val int
     Prev *Node
     Next *Node
     Child *Node
 }

func flatten2(root *Node) *Node  {
	if root == nil {
		return root
	}

	cur := root
	tail := root
	stack := make([]*Node,0)

	for cur != nil || len(stack) != 0 {
		if cur == nil {
			cur = stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			tail.Next = cur
			cur.Prev = tail
		}

		if cur.Child == nil {
			tail = cur
			cur = cur.Next
			continue
		}

		if cur.Next != nil {
			stack = append(stack,cur.Next)
		}
		cur.Child.Prev = cur
		cur.Next = cur.Child
		cur.Child = nil
		tail = cur
		cur = cur.Next
	}

	return root
}

func flatten(root *Node) *Node {
	if root == nil {
		return nil
	}

	newChild := flatten(root.Child)
	newNext := flatten(root.Next)

	if newChild == nil && newNext == nil {
		return root
	}

	if newChild == nil {
		return root
	}

	if newNext == nil {
		root.Next = newChild
		newChild.Prev = root
		root.Child = nil
		return root
	}

	tail := newChild

	for tail.Next != nil {
		tail = tail.Next
	}

	root.Next = newChild
	newChild.Prev = root
	root.Child = nil
	tail.Next = newNext
	newChild.Prev = tail

	return root
}

func main() {
	node1 := &Node{
		Val:   1,
		Prev:  nil,
		Next:  nil,
		Child: nil,
	}

	node2 := &Node{
		Val:   2,
		Prev:  nil,
		Next:  nil,
		Child: nil,
	}
	node1.Next  = node2
	node2.Prev = node1

	node3 := &Node{
		Val:   3,
		Prev:  nil,
		Next:  nil,
		Child: nil,
	}

	node2.Next  = node3
	node3.Prev = node2

	node11 := &Node{
		Val:   11,
		Prev:  nil,
		Next:  nil,
		Child: nil,
	}

	node22 := &Node{
		Val:   22,
		Prev:  nil,
		Next:  nil,
		Child: nil,
	}
	node11.Next  = node22
	node22.Prev = node11

	node33 := &Node{
		Val:   33,
		Prev:  nil,
		Next:  nil,
		Child: nil,
	}
	node22.Next  = node33
	node33.Prev = node22

	node1.Child = node11

	out := flatten2(node1)
	for out != nil {
		fmt.Println(out)
		out = out.Next
	}






}
