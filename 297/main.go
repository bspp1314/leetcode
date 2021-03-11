package main

import (
	"fmt"
	"strconv"
	"strings"
)

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

type Codec struct {

}

func Constructor() Codec {
	return Codec{}
}

// Serializes a tree to a single string.
func (this *Codec) serialize(root *TreeNode) string {
	if root == nil {
		return ""
	}

	q := []*TreeNode{root}
	var res []string
	for len(q) != 0 {
		node := q[0]
		q = q[1:]
		if node != nil {
			res = append(res, strconv.Itoa(node.Val))
			q = append(q, node.Left)
			q = append(q, node.Right)
		} else {
			res = append(res, "null")
		}
	}
	return strings.Join(res, ",")

}

func getValue(a string) int   {
	v, _ := strconv.Atoi(a)
	return v
}

// Deserializes your encoded data to tree.
func (this *Codec) deserialize(data string) *TreeNode {
	if len(data) == 0 {
		return nil
	}
	nodes := strings.Split(data,",")
	if len(nodes) == 0 {
		return nil
	}

	if nodes[0] == "null" {
		return nil
	}

	root := &TreeNode{
		Val:   getValue(nodes[0]),
		Left:  nil,
		Right: nil,
	}

	q := []*TreeNode{root}
	index := 1

	for index < len(nodes) {
		node := q[0]
		q = q[1:]
		leftVal := nodes[index]
		rightVal := nodes[index+1]
		if leftVal != "null" {
			v, _ := strconv.Atoi(leftVal)
			leftNode := &TreeNode{Val: v}
			node.Left = leftNode
			q = append(q, leftNode)
		}
		if rightVal != "null" {
			v, _ := strconv.Atoi(rightVal)
			rightNode := &TreeNode{Val: v}
			node.Right = rightNode
			q = append(q, rightNode)
		}
		index += 2
	}
	return root






}


func main() {
	//node := &TreeNode{
	//	Val:   5,
	//	Left:  &TreeNode{
	//		Val:   0,
	//		Left:  nil,
	//		Right: &TreeNode{
	//			Val:   4,
	//			Left:  nil,
	//			Right: nil,
	//		},
	//	},
	//	Right: nil,
	//}

	c := &Codec{}
	node := c.deserialize("")
	fmt.Println(node)
	//out := c.serialize(node)
	//fmt.Println(out)


}

