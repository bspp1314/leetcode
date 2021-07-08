package main
type ListNode struct {
	Val  int
	Next *ListNode
}


func hasCycle(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return false
	}

	left := head
	right := head
	for {
		left = left.Next
		right = right.Next
		if right == nil {
			return false
		}

		right = right.Next
		if right == nil {
			return false
		}

		if left == right {
			break
		}
	}

	return true
}
func main() {

}
