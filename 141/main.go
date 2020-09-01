package main
type ListNode struct {
	Val  int
	Next *ListNode
}

func hasCycle2(head *ListNode) bool {
	if head == nil {
		return false
	}

	address := make(map[*ListNode]bool)
	for head != nil {
		if address[head] {
			return true
		}
		address[head] = true
		head = head.Next

	}
	return false
}

func hasCycle(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return false
	}

	slow := head
	fast := head
	for {
		if fast == nil {
			return false
		}
		slow = slow.Next
		fast = fast.Next
		if fast == nil {
			return false
		}
		fast = fast.Next

		if slow == fast {
			break
		}
	}

	return true
}
func main() {

}
