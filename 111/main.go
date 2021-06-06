package main
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}


func Min(a,b int) int {
	if a < b {
		return a
	}

	return b
}



func minDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	if root.Right == nil && root.Left == nil {
		return 1
	}

	if root.Right == nil {
		return 1 + minDepth(root.Left)
	}

	if root.Left == nil {
		return 1 + minDepth(root.Right)
	}

	return 1 + Min(minDepth(root.Left),minDepth(root.Right))
}

func minDepth2(root *TreeNode) int {
	if root == nil {
		return 0
	}

	currentLevel := []*TreeNode{root}
	level := 1

	for {
		var nextLevel []*TreeNode
		for i:=0;i<len(currentLevel);i++ {
			node := currentLevel[i]
			if node.Left == nil && node.Right == nil {
				return level
			}

			if node.Left != nil {
				nextLevel = append(nextLevel,node.Left)
			}

			if node.Right != nil {
				nextLevel = append(nextLevel,node.Right)
			}
		}
		level++
		currentLevel = nextLevel
	}


}