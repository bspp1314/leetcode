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

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}

	if l2 == nil {
		return l1
	}

	cur1 := l1
	cur2 := l2

	v1 := make([]*ListNode, 0)
	v2 := make([]*ListNode, 0)

	for cur1 != nil || cur2 != nil {
		if cur1 != nil {
			v1 = append(v1, cur1)
			cur1 = cur1.Next
		}

		if cur2 != nil {
			v2 = append(v2, cur2)
			cur2 = cur2.Next
		}
	}


	if len(v1) < len(v2) {
		v2, v1 = v1, v2
		l1, l2 = l2, l1
	}

	carryValue := 0

	v1Len := len(v1)
	v2Len := len(v2)
	for i:=0;i< v2Len;i++ {
		val := v1[v1Len-i-1].Val + v2[v2Len-i-1].Val + carryValue
		v1[v1Len-i-1].Val = val % 10
		carryValue = val / 10
	}

	for i:= v2Len;i < v1Len;i++ {
		val := v1[v1Len-i-1].Val + carryValue
		v1[v1Len-i-1].Val = val % 10
		carryValue = val / 10
	}

	if carryValue > 0 {
		node := &ListNode{Val: carryValue, Next: l1}
		l1 = node
	}

	return l1
}

func main() {
	l1 := ArrayToListNode([]int{2, 9, 9, 9})
	l2 := ArrayToListNode([]int{1, 1, 9})
	//PrintList(l1)
	//PrintList(l2)
	out := addTwoNumbers(l1, l2)
	PrintList(out)
}

func PrintList(list *ListNode) {
	for list != nil {
		fmt.Printf(" %d ", list.Val)
		list = list.Next
	}
	fmt.Println()
}
