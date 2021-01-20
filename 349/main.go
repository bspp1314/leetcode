package main

import (
	"fmt"
)

//[2,-1,1,2,2]  0->2->3->0
//[-1,2] 1->1->1->>
// [-10,1,3]  0->3->1->2-1->2

func NextIndex(nums []int,index,l int) int  {
	nextStep := nums[index] % l

	if nextStep <  0 {
		nextStep = l + nextStep
	}


	return (nextStep + index) % l

}



// 没有新增
// total  10
// realTotal 10
// new 0


// 新增的当天
// total 12
// realTotal 10
// new 2

// 新增的后一天
// total 14
// realTotal 12
// new 2

func circularArrayLoop(nums []int) bool {
	if len(nums) <= 1 {
		return false
	}

	slow := 0
	fast:= 0
	l := len(nums)

	for i := 0; i < len(nums); i++ {
		slow = NextIndex(nums,slow,l)
		fast = NextIndex(nums,fast,l)
		fast = NextIndex(nums,fast,l)
		if slow == fast {
			break
		}
	}


	if slow ==  NextIndex(nums,fast,l) {
		return false
	}

	//3 1 2
	// next is
	slow = NextIndex(nums,slow,l)
	for slow != fast {
		next := NextIndex(nums,slow,l)
		if nums[slow] * nums[next]  < 0 {
			return false
		}

		slow = next
	}

	return true
}

func intersection(nums1 []int, nums2 []int) []int {
	if len(nums2) == 0 || len(nums2) == 0 {
		return []int{}
	}
	m1 := make(map[int]int)
	for i :=0;i<len(nums1);i++ {
		m1[nums1[i]]++
	}

	res := make([]int,0)
	m2 := make(map[int]int,0)



	for i :=0;i<len(nums2);i++ {
		if m1[nums2[i]] > m2[nums2[i]]  {
			res = append(res,nums2[i])
			m2[nums2[i]]++
		}
	}

	return res
}

func main() {
	//fmt.Println(circularArrayLoop([]int{2,-1,1,2,2}))
	//fmt.Println(circularArrayLoop([]int{-1,2}))
	//fmt.Println(circularArrayLoop([]int{-2,1,-1,-2,-2}))
	//fmt.Println(circularArrayLoop([]int{2,-1,1,-2,-2}))
	//fmt.Println(circularArrayLoop([]int{-1,-2,-3,-4,-5}))
	fmt.Println(circularArrayLoop([]int{3,1,2}))
	fmt.Println(circularArrayLoop([]int{4,1,2,3}))
}


