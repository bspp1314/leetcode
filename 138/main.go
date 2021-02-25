package main

type Node struct {
	Val    int
	Next   *Node
	Random *Node
}



var Visit 	 = make(map[*Node]*Node)
func copyRandomList(head *Node) *Node {
	if head == nil {
		return nil
	}

	if v,ok := Visit[head];ok {
		return v
	}

	if head.Next == nil || head.Random == nil {
		return &Node{Val: head.Val}
	}

	copyNode := &Node{
		Val:    head.Val,
		Next:   nil,
		Random: nil,
	}
	Visit[head] = copyNode
	if head.Next != nil {
		copyNode.Next = copyRandomList(head.Next)
	}

	if head.Random != nil {
		copyNode.Random = copyRandomList(head.Random)
	}

	return copyNode

}
