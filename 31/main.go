package main

import "fmt"

func reverse(nums []int,left int,right int) {
	for left < right {
		nums[left],nums[right]  = nums[right],nums[left]
		left++
		right--
	}
}

func nextPermutation(nums []int) []int {

	k := 0
	for i:=len(nums)-1;i > 0 ;i-- {
		if nums[i] <= nums[i-1] {
			continue
		}

		for j:= len(nums)-1;j>=i;j-- {
			if nums[j] > nums[i-1] {
				nums[j],nums[i-1] = nums[i-1],nums[j]
				break
			}
		}
		k = i
		break
	}

	reverse(nums,k,len(nums)-1)
	return nums

}
func main() {
	fmt.Println(nextPermutation([]int{-1,2,10,30,4,3,2,1}))
}