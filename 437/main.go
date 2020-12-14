package main

import (
	"fmt"
)

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func pathSum(root *TreeNode, sum int) int {
	// 初始化用于记录路径和
	dic := make(map[int]int)
	// 路径为0的1条
	dic[0] = 1

	return pathSumHelp(root,0,sum,dic)
}

func pathSumHelp(root *TreeNode, nowSum, sum int,dic map[int]int) int   {
	if root == nil {
		return 0
	}

	nowSum += root.Val //当前路径上的和
	//(nowSum - sum) 表示从root节点到当前节点以上节点可能存在的和
	// 比如  sum = 10
	//         8
	//       /
	//     9
	//    /
	//   9
	//  /
	// 1
	// 到达节点1的时候 dic[0:1,8:0,26:0]
	// nowSum = 27 - 10 =  17
	// nowSum-sum 就是对前缀节点进行移除的一次操作
	times := dic[nowSum-sum]
	dic[nowSum]++


	times += pathSumHelp(root.Left,nowSum,sum,dic)
	times += pathSumHelp(root.Right,nowSum,sum,dic)
	// 回溯，是为了让底层的值不影响到上层的值
	dic[nowSum]--

	return times

}


func main() {
	root2 := &TreeNode{
		Val:   8,
		Left:  &TreeNode{
			Val:   9,
			Left:  &TreeNode{
				Val:   9,
				Left:  &TreeNode{
					Val:   1,
					Left:  nil,
					Right: nil,
				},
				Right: nil,
			},
			Right: nil,
		},
		Right: nil,
	}
	fmt.Println(pathSum(root2,10))
}
