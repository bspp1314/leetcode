package main

import "fmt"

func subsets(nums []int) [][]int {
	if len(nums) == 0 {
		return [][]int{[]int{}}
	}

	numsLastIndex := len(nums)-1
	sub := subsets(nums[:numsLastIndex])

	subLen := len(sub)

	for i := 0; i < subLen; i++ {
		tem := make([]int,len(sub[i]))
		copy(tem,sub[i])
		sub = append(sub,append(tem,nums[numsLastIndex]))
	}


	return sub

}

func main() {
	fmt.Println(subsets([]int{9,0,3,5,7}))
}