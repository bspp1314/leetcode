package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
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

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}

	if l2 == nil {
		return l1
	}

	var head *ListNode
	var tail *ListNode

	if l1.Val < l2.Val {
		head = l1
		l1 = l1.Next
		head.Next = nil
	}else{
		head = l2
		l2 = l2.Next
		head.Next = nil
	}

	tail = head

	for l1 != nil && l2 != nil {
		if l1.Val < l2.Val {
			n := l1.Next
			l1.Next = nil
			tail.Next = l1
			tail = tail.Next
			l1 = n
		}else{
			n := l2.Next
			l2.Next = nil
			tail.Next = l2
			tail = tail.Next
			l2 = n
		}
	}

	if  l1 != nil {
		tail.Next = l1
	}

	if l2 != nil {
		tail.Next = l2
	}

	return head


}




func mergeKLists(lists []*ListNode) *ListNode {
	if len(lists) == 0 {
		return nil
	}

	if len(lists) == 1 {
		return lists[0]
	}

	if len(lists) == 2 {
		return mergeTwoLists(lists[0],lists[1])
	}

	mid := (len(lists) -1)/2
	l1 := mergeKLists(lists[0:mid])
	l2 := mergeKLists(lists[mid:len(lists)])

	return mergeTwoLists(l1,l2)
}

func main() {
	k := [][]int{{1,4,5},{1,3,4},{2,6}}
	var input []*ListNode

	for i := 0; i < len(k); i++ {
		input = append(input,CreateList(k[i]))
	}

	PrintList(mergeKLists(input))
}