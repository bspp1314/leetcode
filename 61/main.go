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

func rotateRight(head *ListNode, k int) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	if k == 0 {
		return head
	}

	listMap := make(map[int]*ListNode)

	tail := head
	index := 0

	for tail != nil {
		listMap[index] = tail
		tail = tail.Next
		index++
	}

	k = k % index

	if k == 0 {
		return head
	}

	listMap[index-1].Next = head
	listMap[index-k-1].Next = nil

	return listMap[index-k]
}

func main() {
	arr := []int{1, 2, 3, 4, 5}
	list := ArrayToListNode(arr)
	PrintList(list)
	newList := rotateRight(list, 2)
	PrintList(newList)
}
