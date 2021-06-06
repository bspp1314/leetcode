package main

type Node struct {
	Val   int
	Left  *Node
	Right *Node
	Next  *Node
}

func connect(root *Node) *Node {
	if root == nil {
		return nil
	}

	if root.Left == nil && root.Right == nil {
		return root
	}

	nowLevel := root

	for {
		var nextLevelHead *Node
		var nextLevelTail *Node

		for nowLevel != nil {
			if nowLevel.Left != nil {
				if nextLevelHead == nil {
					nextLevelHead =nowLevel.Left
					nextLevelTail = nowLevel.Left
				}else{
					nextLevelTail.Next = nowLevel.Left
					nextLevelTail = nextLevelTail.Next
				}
			}

			if nowLevel.Right != nil {
				if nextLevelHead == nil {
					nextLevelHead =nowLevel.Right
					nextLevelTail = nowLevel.Right
				}else{
					nextLevelTail.Next = nowLevel.Right
					nextLevelTail = nextLevelTail.Next
				}
			}
		}

		if nextLevelHead == nil {
			break
		}

		nowLevel = nextLevelHead
	}

	return root
}
