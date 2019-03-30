package main

import (
	"fmt"
	"log"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func PrintlnList(head *ListNode) {
	for head != nil {
		fmt.Printf("%d ", head.Val)
		head = head.Next
	}
	fmt.Println()
}
func NewLists(lists []int) *ListNode {
	if len(lists) < 0 {
		return nil
	}

	head := &ListNode{
		Val:  lists[0],
		Next: nil,
	}
	tmp := head
	for i := 1; i < len(lists); i++ {
		AddNode(tmp, lists[i])
		tmp = tmp.Next
	}
	return head

}
func AddNode(head *ListNode, val int) {
	if head == nil {
		return
	}
	head.Next = &ListNode{
		Val:  val,
		Next: nil,
	}
}
func reverseList(head *ListNode) *ListNode {
	var resultHead *ListNode
	var tmp *ListNode
	for head != nil {
		if resultHead == nil {
			resultHead = head
			head = head.Next
			resultHead.Next = nil
		} else {
			tmp = head
			head = head.Next
			tmp.Next = resultHead
			resultHead = tmp
		}
	}
	return resultHead
}
func reverseKGroup(head *ListNode, k int) *ListNode {
	var resultHead *ListNode
	var resultTail *ListNode
	var newKHead *ListNode
	var newKTail *ListNode
	var tmp *ListNode

	if k == 0 || k == 1 {
		return reverseList(head)
	}
	i := 0
	for head != nil {
		if i%k == 0 {
			if resultHead == nil {
				if i/k == 1 {
					resultHead = newKHead
				}
			} else {
				resultTail.Next = newKHead
			}
			resultTail = newKTail

			newKHead = head
			newKTail = newKHead
			head = head.Next
			newKHead.Next = nil
		} else {
			tmp = head
			head = head.Next
			tmp.Next = newKHead
			newKHead = tmp
		}
		i++
	}

	if i%k != 0 {
		if resultHead == nil {
			resultHead = reverseList(newKHead)
		} else {
			resultTail.Next = reverseList(newKHead)
		}
	} else {
		if resultHead == nil {
			if i/k == 1 {
				resultHead = newKHead
			}
		} else {
			resultTail.Next = newKHead
		}
		resultTail = newKTail
	}
	return resultHead
}

func reverseKGroup2(head *ListNode, k int) *ListNode {
	tail := head
	for i := 0; i < k; i++ {
		if tail == nil {
			return head
		}
		tail = tail.Next
	}
	tail = head
	for i := 0; i < k-1; i++ {
		tmp := tail.Next
		tail.Next, head, tmp.Next = tmp.Next, tmp, head
	}
	tail.Next = reverseKGroup2(tail.Next, k)
	return head
}
func main() {
	lists := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	head := NewLists(lists)
	//reverseKGroup(head, 3)

	//newHead := reverseKGroup(head, 5)
	//PrintlnList(newHead)

	head = NewLists(lists)
	newHead := reverseKGroup2(head, 9)
	log.Println(newHead)
	PrintlnList(newHead)
}
