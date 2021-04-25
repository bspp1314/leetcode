package main

import "fmt"

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

func PrintList(l *ListNode)  {
	for l != nil {
		fmt.Printf("%d ",l.Val)
		l = l.Next
	}

	fmt.Println()
}

func oddEvenList2(head *ListNode) *ListNode {
	if head == nil {
		return head
	}
	rightHead := head.Next
	left := head
	right := rightHead
	for right != nil && right.Next != nil {
		left.Next = right.Next
		left = left.Next
		right.Next = left.Next
		right = right.Next
	}
	left.Next = rightHead
	return head
}




func oddEvenList(head *ListNode) *ListNode {
	if head == nil {
		return head
	}

	if head.Next == nil {
		return head
	}

	if head.Next.Next == nil {
		return head
	}


	tem := head.Next.Next
	leftHead := head
	rightHead := head.Next
	left := leftHead
	right := rightHead
	left.Next = nil
	right.Next = nil
	i := true

	for tem != nil {
		if i  {
			left.Next = tem
			left = left.Next
			i = false
			tem = tem.Next
			left.Next = nil
		}else{
			right.Next = tem
			right = right.Next
			i = true
			tem = tem.Next
			right.Next = nil
		}
	}

	left.Next  = rightHead
	return head
}
func main() {
	list := ArrayToListNode([]int{1,2,3,4,5,6})
	out := oddEvenList(list)
	PrintList(out)
}
