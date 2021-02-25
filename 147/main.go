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



func insertionSortList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	fakeHead := &ListNode{Next: head}

	lastSorted := head
	curr := head.Next
	for curr != nil  {
		if lastSorted.Val <= curr.Val {
			lastSorted = lastSorted.Next
		} else {
			prev := fakeHead
			for prev.Next.Val <= curr.Val {
				prev = prev.Next
			}

			lastSorted.Next = curr.Next
			curr.Next = prev.Next
			prev.Next = curr

		}

		curr = lastSorted.Next
	}


	return fakeHead.Next
}

func main() {
	head := []int{9,1}
	list := CreateList(head)
	PrintList(insertionSortList(list))
}
