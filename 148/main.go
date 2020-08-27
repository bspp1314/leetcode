package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func sortList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	sortListHelp(head,nil)
	return head
}

func PrintList(list *ListNode) {
	for list != nil {
		fmt.Printf(" %d ", list.Val)
		list = list.Next
	}
	fmt.Println()
}

func ArrayToListNode(arr []int) (*ListNode, *ListNode) {
	if len(arr) == 0 {
		return nil, nil
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

	return head, tail

}
func sortListHelp(head *ListNode, tail *ListNode) {
	if head == tail {
		return
	}

	slow := head
	fast := head
	v := head.Val

	for fast != tail {
		if fast.Val >= v {
			fast = fast.Next
		} else {
			if slow == fast {
				fast.Val = fast.Next.Val
				fast = fast.Next
				fast.Val = v
				slow = fast
			} else {
				slow.Val = fast.Val
				fast.Val = slow.Next.Val
				slow = slow.Next
				slow.Val = v
				fast = fast.Next
			}
		}
	}

	sortListHelp(head, slow)
	sortListHelp(slow.Next, tail)
}

func merge(left, right *ListNode) *ListNode {
	tmp := &ListNode{-1, nil}
	h := tmp
	for left != nil && right != nil {
		if left.Val < right.Val {
			tmp.Next = left
			tmp = left
			left = left.Next
		}   else {
			tmp.Next = right
			tmp = right
			right = right.Next
		}
	}
	if left != nil {
		tmp.Next = left
	} else if right != nil {
		tmp.Next = right
	}
	return h.Next
}


func main() {
	head, _ := ArrayToListNode([]int{2,1})
	PrintList(head)

}
