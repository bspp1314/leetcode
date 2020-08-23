package main

import (
	"fmt"
)

type ListNode struct {
     Val int
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

	for i:=1;i<len(arr);i++ {
		new := &ListNode{
			Val:  arr[i],
			Next: nil,
		}

		tail.Next = new
		tail = tail.Next
	}

	return head

}

func PrintList(list *ListNode)  {
	for list != nil   {
		fmt.Printf(" %d ",list.Val)
		list = list.Next
	}
	fmt.Println()
}

func rotateRight(head *ListNode, k int) *ListNode {
	if head == nil || k == 0 {
		return head
	}

	l := 0
	tail := head
	address := make(map[int]*ListNode)
	for tail != nil {
		address[l] = tail
		tail = tail.Next
		l++
	}

	k = k % l
	if k == 0 {
		return  head
	}
	newHead := address[l-k]
	newTail := address[l-k-1]
	newTail.Next = nil
	address[l-1].Next = head

	return newHead
}

func rotateRight2(head *ListNode, k int) *ListNode {
	if head == nil || k == 0 {
		return head
	}

	l := 0
	tail := head
	for tail.Next != nil  {
		tail = tail.Next
		l++
	}

	fmt.Println(tail)
	k = k % (l+1)
	if k == 0 {
		return  head
	}

	tail.Next = head

	newTail := head
	for i:= 0;i<l-k;i++ {
		newTail = newTail.Next
	}

	newHead := newTail.Next
	newTail.Next = nil

	return newHead
}

func main()  {
	arr :=[]int{1,2,3,4,5,6}
	list := ArrayToListNode(arr)
	PrintList(list)
	newList := rotateRight(list,5)
	PrintList(newList)
	list2 := ArrayToListNode(arr)
	newList2 := rotateRight2(list2,5)
	PrintList(newList2)
}

