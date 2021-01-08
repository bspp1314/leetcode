package main

import "fmt"

type ListNode struct {
	Next *ListNode
	Val int
}

func PrintList(l *ListNode)  {
	for l != nil {
		fmt.Printf("%d ",l.Val)
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
	var newHead *ListNode

	for head != nil {
		t := head
		head = head.Next
		t.Next = newHead
		newHead = t
	}
	return newHead

}

func middleNode(head *ListNode) (*ListNode,*ListNode) {
	slow:= head
	fast := head
	var slow2 *ListNode


	for fast != nil {
		fast = fast.Next
		if fast == nil {
			break
		}

		fast = fast.Next
		slow2 = slow
		slow  = slow.Next

	}

	return slow,slow2
}

func reorderList(head *ListNode)  {
	if head == nil  || head.Next == nil || head.Next.Next == nil {
		return
	}

	s1,s2 := middleNode(head)
	if s2 != nil {
		s2.Next = nil
	}

	s1 = reverseList(s1)
	PrintList(head)
	PrintList(s1)

	var newHead *ListNode
	var cur *ListNode

	for head != nil {
		if newHead == nil {
			newHead = head
			cur = head
			head = head.Next
		}else{
			t := head
			head = head.Next
			cur.Next = t
			cur = t
			cur.Next = nil
		}

		t := s1
		s1 = s1.Next
		cur.Next = t
		cur = t
		cur.Next = nil
	}

	if s1 != nil {
		t := s1
		s1 = s1.Next
		cur.Next = t
		cur = t
		cur.Next = nil
	}

	PrintList(newHead)
	head = newHead

}

func main() {
	l := CreateList([]int{1,2,3,4,5,6,7})
	reorderList(l)
}
