package main

import "fmt"

/*
给定两个非空链表来表示两个非负整数。位数按照逆序方式存储，它们的每个节点只存储单个数字。将两数相加返回一个新的链表。

你可以假设除了数字 0 之外，这两个数字都不会以零开头。

示例：

输入：(2 -> 4 -> 3) + (5 -> 6 -> 4)
输出：7 -> 0 -> 8
原因：342 + 465 = 807

*/
/**
 * Definition for singly-linked list.
 */
type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	low := 0
	high := 0
	var head *ListNode
	var tail *ListNode
	for l1 != nil || l2 != nil {
		total := 0


		if l1 != nil {
			total += l1.Val
			l1 = l1.Next
		}

		if l2 != nil {
			total+= l2.Val
			l2 = l2.Next
		}

		total += low
		high = total % 10
		low  = total / 10
		if head == nil {
			head = &ListNode{
				Val:high,
			}
			tail = head
		}else{
			tail.Next = &ListNode{Val:high}
			tail = tail.Next
		}
	}

	if low != 0 {
		tail.Next = &ListNode{Val:low}
	}

	return head

}

func main() {
	node1 := &ListNode{Val: 7}
	node2 := &ListNode{
		Val:  8,
		Next: node1,
	}
	node3 := &ListNode{
		Val:  8,
		Next: node2,
	}
	node4 := &ListNode{Val: 5}
	node5 := &ListNode{
		Val:  8,
		Next: node4,
	}
	node6 := &ListNode{
		Val:  6,
		Next: node5,
	}

	result := addTwoNumbers(node3, node6)
	for result != nil {
		fmt.Println("Val", result.Val)
		result = result.Next
	}
}
