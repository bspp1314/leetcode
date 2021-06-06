package main



  type ListNode struct {
      Val int
      Next *ListNode
  }
  type TreeNode struct {
      Val int
      Left *TreeNode
      Right *TreeNode
  }

func sortedListToBST(head *ListNode) *TreeNode {
    if head == nil {
        return nil
    }

    nums := make([]int,0)
    root := head

    for root != nil {
        nums = append(nums,root.Val)
        root = root.Next
    }

    return sortedArrayToBST(nums)


}

func sortedArrayToBST(nums []int) *TreeNode {
    if len(nums) == 0 {
        return nil
    }

    if len(nums) == 1 {
        return &TreeNode{Val:nums[0]}
    }

    tree := &TreeNode{}
    mid := len(nums) / 2

    // get left
    tree.Val = nums[mid]
    tree.Left  = sortedArrayToBST(nums[:mid])
    tree.Right = sortedArrayToBST(nums[mid+1:])


    return tree
}

func main() {

}
