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
			if nextLevelHead == nil {
				if nowLevel.Left != nil && nowLevel.Right != nil {
					nextLevelHead = nowLevel.Left
					nextLevelTail = nowLevel.Right
					nextLevelHead.Next = nextLevelTail
				}


			} else {
				if nowLevel.Left != nil && nowLevel.Right != nil {
					nextLevelTail.Next = nowLevel.Left
					nextLevelTail = nextLevelTail.Next
					nextLevelTail.Next = nowLevel.Right
					nextLevelTail = nextLevelTail.Next
				}
			}

			nowLevel = nowLevel.Next
		}

		if nextLevelHead == nil {
			break
		}

		nowLevel = nextLevelHead
	}

	return root
}
