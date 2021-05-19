package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func PrintList(l *ListNode) {
	for l != nil {
		fmt.Printf("%d ", l.Val)
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

func swapPairs(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	t := head.Next
	head.Next = t.Next
	t.Next = head
	head = t

	newHead := swapPairs(head.Next.Next)
	head.Next.Next = newHead


	return head

}

func main() {
	list := CreateList([]int{1,2,3,4,5,6})
	list = swapPairs(list)
	PrintList(list)
}
