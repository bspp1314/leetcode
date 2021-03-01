package main

import "fmt"

type ListNode struct {
      Val int
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


func removeElements(head *ListNode, val int) *ListNode {
	if head == nil {
		return nil
	}

	if head.Val == val {
		return removeElements(head.Next,val)
	}

	pre  := head
	next := head.Next

	for next != nil {
		if next.Val ==  val {
			pre.Next = next.Next
			next = pre.Next
		}else{
			pre = next
			next = next.Next
		}
	}

	return head
}

func main() {
	list := CreateList([]int{6,2,6,4,6})
	PrintList(list)
	PrintList(removeElements(list,2))

}
