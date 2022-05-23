package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func generateTrees(n int) []*TreeNode {
	if n == 0 {
		return nil
	}
	nMap := make(map[[2]int][]*TreeNode)
	return helper(1, n, nMap)
}

func helper(start, end int, nMap map[[2]int][]*TreeNode) []*TreeNode {
	if start > end {
		trees := []*TreeNode{nil}
		nMap[[2]int{start, end}] = trees
		return trees
	}
	var allTrees []*TreeNode
	// 枚举可行根节点
	for i := start; i <= end; i++ {
		// 获得所有可行的左子树集合
		var leftTrees []*TreeNode
		var rightTrees []*TreeNode

		if v, ok := nMap[[2]int{start, i - 1}]; ok {
			leftTrees = v
		} else {
			leftTrees = helper(start, i-1, nMap)
		}

		if v, ok := nMap[[2]int{i + 1, end}]; ok {
			rightTrees = v
		} else {
			rightTrees = helper(i+1, end, nMap)
		}

		// 从左子树集合中选出一棵左子树，从右子树集合中选出一棵右子树，拼接到根节点上
		for _, left := range leftTrees {
			for _, right := range rightTrees {
				currTree := &TreeNode{i, nil, nil}
				currTree.Left = left
				currTree.Right = right
				allTrees = append(allTrees, currTree)
			}
		}
	}
	nMap[[2]int{start, end}] = allTrees
	return allTrees
}

func main() {
	nMap := make(map[[2]int]int)
	a := [2]int{}
	nMap[a] = 3
	b := [2]int{}
	fmt.Println(nMap[b])
}
