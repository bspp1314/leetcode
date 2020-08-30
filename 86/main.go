package main

import (
	"fmt"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func ArrayToListNode(arr []int) *ListNode {
	if len(arr) == 0 {
		return nil
	}

	head := &ListNode{
		Val:  arr[0],
		Next: nil,
	}

	tail := head

	for i := 1; i < len(arr); i++ {
		new := &ListNode{
			Val:  arr[i],
			Next: nil,
		}

		tail.Next = new
		tail = tail.Next
	}

	return head

}

func PrintList(list *ListNode) {
	for list != nil {
		fmt.Printf(" %d ", list.Val)
		list = list.Next
	}
	fmt.Println()
}
func partition(head *ListNode, x int) *ListNode {
	//  4-> 4 -> 3 -> 2 ->5 ->2
	if head == nil || head.Next == nil {
		return head
	}

	var slow *ListNode
	var fast2 *ListNode
	fast := head.Next
	fast2 = head
	if fast2.Val < x {
		slow = head
	}

	for fast != nil {
		if fast.Val >= x {
			fast2 = fast
			fast = fast.Next
			continue
		}

		if fast2.Val < x {
			fast2 = fast
			fast = fast2.Next
			slow = fast2
			continue
		}

		tem := fast
		fast2.Next = fast.Next
		fast = fast2.Next
		tem.Next = nil

		if slow == nil {
			tem.Next = head
			slow = tem
			head = slow
		} else {
			tem.Next = slow.Next
			slow.Next = tem
			slow = slow.Next
		}
	}

	return head
}

func partition2(head *ListNode, x int) *ListNode {
	p1Head := &ListNode{}
	p2Head := &ListNode{}

	p1 := p1Head
	p2 := p2Head

	for head != nil {
		if head.Val < x {
			p1.Next = head
			head = head.Next
			p1 = p1.Next
			p1.Next = nil
		} else {
			p2.Next = head
			head = head.Next
			p1 = p1.Next
			p1.Next = nil
		}
	}

	p1.Next = p2Head.Next
	return p1Head.Next
}

func main() {
	head := ArrayToListNode([]int{6, 5, 4, 3, 2, 1})
	head = partition(head, 4)
	PrintList(head)
}
