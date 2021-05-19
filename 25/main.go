package main

import (
	"fmt"
)

type ListNode struct {
	Val  int
	Next *ListNode
}


func PrintList(l *ListNode) {
	for l != nil {
		fmt.Printf("%d ", l.Val)
		l = l.Next
	}

	fmt.Println()
}

func CreateList(nodes []int) *ListNode {
	if len(nodes) == 0 {
		return nil
	}

	head := &ListNode{Val: nodes[0]}
	cur := head

	for i := 1; i < len(nodes); i++ {
		cur.Next = &ListNode{
			Next: nil,
			Val:  nodes[i],
		}

		cur = cur.Next
	}

	return head
}

func reverseList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	res := reverseList(head.Next)
	t := head.Next
	head.Next = nil
	t.Next = head
	return res
}

func reverseKGroup(head *ListNode, k int) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	tail := head
	t := k
	for t > 1 && tail != nil {
		t --
		tail = tail.Next
	}

	if t > 1 || tail == nil {
		return head
	}


	v := reverseKGroup(tail.Next,k)
	tail.Next = nil
	newHead := reverseList(head)
	head.Next = v
	return newHead
}

func main() {
	lists := []int{1, 2, 3, 4, 5}
	list := CreateList(lists)
	out := reverseKGroup(list,3)
	PrintList(out)
}
