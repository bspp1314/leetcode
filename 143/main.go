package main

type ListNode struct {
	Next *ListNode
	Val int
}

// 1 2 3 4 5 6 7

// slow 1
// fast 1

// slow 2
// fast 3 4

// slow 3
// fast 5 6


// 1 2 3 4 5 6 7 8


// slow 1
// fast 1

// slow 2
// fast 3 4

// slow 3
// fast 5 6

// slow 4
// fast 7 8


func middleNode(head *ListNode) *ListNode  {
	slow,fast := head,head

	for fast.Next !=nil && fast.Next.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}

	return slow
}

func mergerList(l1,l2 *ListNode) {
	var l1Temp  *ListNode
	var l2Temp *ListNode

	for l1 != nil && l2 != nil {
		// 1->2->3
		// 6->5-4


		l1Temp = l1.Next
		l2Temp = l2.Next

		// 1->6->5-4
		// 2->3
		// l1 = 5->4
		l1.Next = l2
		l1 = l1Temp
		// 1->6->2-3
		// l2 = 5->4
		// l1 = 1->6->2->3
		l2.Next = l1
		l2 = l2Temp
	}
}

func reverseList(head *ListNode) *ListNode  {
	var newHead *ListNode

	for head != nil {
		t := head
		head = head.Next
		t.Next = newHead
		newHead = t
	}

	return newHead
}

func reorderList(head *ListNode)  {
	if head == nil {
		return
	}
	mid := middleNode(head)
	l1 := head
	l2 := mid.Next
	mid.Next = nil
	l2 = reverseList(l2)
	mergerList(l1,l2)
}

func main() {
}
