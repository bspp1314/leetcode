package main


  type TreeNode struct {
      Val int
      Left *TreeNode
      Right *TreeNode
  }

type BSTIterator struct {
	Stack []*TreeNode
	Current *TreeNode

}


func Constructor(root *TreeNode) BSTIterator {
    if root == nil {
        return BSTIterator{}
    }
    res := BSTIterator{
        Stack:   nil,
        Current: nil,
    }
    res.Stack = []*TreeNode{root}
    res.Current = root.Left

    return res
}


func (this *BSTIterator) Next() int {
    for this.Current != nil {
        this.Stack = append(this.Stack,this.Current)
        this.Current = this.Current.Left
    }

    if len(this.Stack) > 0 {
        res := this.Stack[len(this.Stack)-1].Val
        this.Current = this.Stack[len(this.Stack)-1].Right
        this.Stack = this.Stack[:len(this.Stack)-1]
        return res
    }


    return  -1
}


func (this *BSTIterator) HasNext() bool {
    return len(this.Stack) > 0 && this.Current != nil
}