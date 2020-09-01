package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func detectCycle(head *ListNode) *ListNode {
	address := make(map[*ListNode]bool)

	for head != nil {
		_,ok := address[head]
		if ok {
			return head
		}

		address[head] = true
		head = head.Next
	}

	return nil
}
// 1 2 3 2 4 2
func detectCycle2(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return  nil
	}

	slow := head
	fast := head
	hasCycle := false
	for slow.Next != nil && fast.Next.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
		if slow == fast {
			hasCycle = true
			break
		}
	}


	if !hasCycle {
		return  nil
	}
	slow = head

	for slow != fast {
		slow = slow.Next
		fast = fast.Next
	}

	return slow
}


func main() {

}
