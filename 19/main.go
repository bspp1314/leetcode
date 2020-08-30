package main

import "fmt"

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


func main() {
	head := ArrayToListNode([]int{1, 2, 3, 4, 5, 6})
	head = removeNthFromEnd(head, 2)
	PrintList(head)
}
