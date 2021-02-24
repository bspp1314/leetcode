package main

import "fmt"

type ListNode struct {
     Val int
     Next *ListNode
 }

// 1->1->1->1
func deleteDuplicates(head *ListNode) *ListNode {
    if head == nil || head.Next == nil {
        return head
    }

    slow := head
    fast := head.Next

    if slow.Val != fast.Val {
        slow.Next = deleteDuplicates(fast)
        return slow
    }


    for fast != nil && slow.Val == fast.Val {
        fast = fast.Next
    }


    return deleteDuplicates(fast)
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
func main() {
    list := CreateList([]int{1,1,1,2})
    PrintList(list)
    PrintList(deleteDuplicates(list))
}
