package main

import "fmt"

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

	ret := reverseList(head.Next)
	tail := head.Next
	head.Next = nil
	tail.Next = head

	return ret
}

func reverseList2(head *ListNode) *ListNode {
	var newHead *ListNode

	for head != nil {
		t := head
		head = head.Next
		t.Next = newHead
		newHead = t
	}
	return newHead

}

func main() {
	node1 := CreateList([]int{1,2,3,4,5})
	node2 := CreateList([]int{1,2,3,4,5})
	node1 = reverseList(node1)
	node2 = reverseList2(node2)
	PrintList(node1)
	PrintList(node2)
}