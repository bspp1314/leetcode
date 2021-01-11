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

func isPalindrome(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return true
	}
	// 1 2 3 2 1
	l1,l2 := middleNode(head)
	l2.Next = nil

	l1 = reverseList(l1)
	for l1 != nil && head != nil {
		if l1.Val != head.Val  {
			return false
		}

		l1 = l1.Next
		head = head.Next
	}

	return true
}

func main() {
	fmt.Println(isPalindrome(ArrayToListNode([]int{1,2,3,2,1})))
}