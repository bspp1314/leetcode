package main

type Node struct {
	Val   int
	Left  *Node
	Right *Node
	Next  *Node
}

func connect2(root *Node) *Node {
	if root == nil {
		return nil
	}

	if root.Left != nil && root.Right != nil {
		root.Left.Next = root.Right
	}

	if root.Left != nil && root.Right == nil {
		root.Left.Next = getNext(root.Next)
	}

	if root.Right != nil {
		root.Right.Next = getNext(root.Next)
	}

	connect(root.Left)
	connect(root.Right)

	return root
}

func connect(root *Node) *Node {
	if root == nil {
		return root
	}

	if root.Left != nil && root.Right != nil {
		root.Left.Next = root.Right
	}

	if root.Left != nil && root.Right == nil {
		root.Left.Next = getNext(root.Next)
	}

	if root.Right != nil {
		root.Right.Next = getNext(root.Next)
	}

	connect(root.Right)
	connect(root.Left)
	return root
}

func getNext(root *Node) *Node {
	if root == nil {
		return nil
	}
	if root.Left != nil {
		return root.Left
	}
	if root.Right != nil {
		return root.Right
	}

	return getNext(root.Next)
}

func getNext2(root *Node) *Node {
	if root == nil {
		return nil
	}

	if root.Left != nil {
		return root.Left
	}

	if root.Right != nil {
		return root.Right
	}

	return getNext2(root.Next)
}




