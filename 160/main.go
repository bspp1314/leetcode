package main

type ListNode struct {
	Val  int
	Next *ListNode
}

// 两个链表相交
// headA 的长度 为 m = a +c
// headB 的长度 为 n = b + c

// 如果 A 和 B 相交
// 如果 a = b.两指针走完 a 个节点就可以达到交点
// 如果 a != b 两指针走完 a+b+c  个节点就可以达到交点，指针1 的路程为 a->c-b,指针 2 的路程为 b->c->a

// 如果 A 和 B 不相交
// m = n,走完 m 路程即可。
// m != n,指针1 的路程为 m->n,指针 2 的路程为 n-m
func getIntersectionNode(headA, headB *ListNode) *ListNode {
	if headA == nil || headB == nil {
		return nil
	}
	pA := headA
	pB := headB

	for pA != pB {
		if pA == nil {
			pA = headB
		} else {
			pA = pA.Next
		}

		if pB == nil {
			pB = headA
		}else{
			pB = pB.Next
		}
	}


	return pA
}
func main() {

}