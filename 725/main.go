package main

import "fmt"

type ListNode struct {
	Val int
	Next *ListNode
}

func splitListToParts(root *ListNode, k int) []*ListNode {
	listNum := make([]int,k)
	head := root

	for head != nil {
		for i:=0;i<k;i++ {
			listNum[i]++
			head = head.Next
			if head == nil {
				break
			}
		}
	}

	resHead := make([]*ListNode,k)
	head = root

	for i:=0;i < k;i++ {
		if listNum[i] == 0 {
			return resHead
		}


		left := head
		right := head
		for listNum[i] != 1  {

			right = right.Next
			listNum[i]--
		}

		head = right.Next
		resHead[i] = left
		right.Next = nil
	}


	return resHead
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



func main() {
	list := ArrayToListNode([]int{1,2,3,4,5,6,7,8,9,10,11,12})
	out := splitListToParts(list,5)

	for i, _ := range out {
		PrintList(out[i])
	}
}
