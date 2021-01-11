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

func reverseBetween(head *ListNode, m int, n int) *ListNode {
	if head == nil || head.Next == nil || m == n {
		return head
	}

	cur := head
	var mHeadBefor *ListNode
	var mHead *ListNode

	i := 1

	for cur != nil {
		if i < m  {
			if (i+1) == m {
				mHeadBefor = cur
			}
			cur = cur.Next
			i++
			continue
		}else if i == m {
			mHead = cur
			break
		}
	}

	var newHead *ListNode
	cur = mHead

	for cur != nil {
		if i > n {
			mHead.Next = cur
			break
		}
		t := cur
		cur = cur.Next
		t.Next = newHead
		newHead = t
		i++
	}

	if mHeadBefor != nil {
		mHeadBefor.Next = newHead
	}else{
		return newHead
	}

	return head

}

func main() {
	PrintList(reverseBetween(ArrayToListNode([]int{1,2,3,4,5}),2,5))
}