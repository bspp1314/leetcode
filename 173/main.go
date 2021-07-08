package main


type TreeNode struct {
      Val int
      Left *TreeNode
      Right *TreeNode
}

type BSTIterator struct {
    Root  *TreeNode
    Stack []*TreeNode
}


func Constructor(root *TreeNode) BSTIterator {
    if root == nil {
        return BSTIterator{}
    }

    return BSTIterator{
        Root:  root,
        Stack: make([]*TreeNode,0),
    }
}


func (this *BSTIterator) Next() int {
    for this.Root != nil {
        this.Stack = append(this.Stack,this.Root)
        this.Root = this.Root.Left
    }

    if len(this.Stack) > 0 {
        n :=this.Stack[len(this.Stack)-1]
        this.Stack = this.Stack[:len(this.Stack)-1]
        this.Root = n.Right

        return n.Val

    }

    return -1
}


func (this *BSTIterator) HasNext() bool {
    return this.Root != nil || len(this.Stack) != 0
}

