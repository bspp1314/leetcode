package main

import (
	"fmt"
)

func productExceptSelf(nums []int) []int {
	if len(nums) == 0 {
		return []int{}
	}

	if len(nums) == 1 {
		return []int{1}
	}
	n := len(nums)
	product := make([]int, n)
	product2 := make([]int, n)
	product[0] = nums[0]
	product2[n-1] = nums[n-1]

	for i := 1; i < n; i++ {
		product[i] = product[i-1] * nums[i]
		product2[n-i-1] = product2[n-i] * nums[n-i-1]
	}

	for i := 0; i < n; i++ {
		if i == 0 {
			nums[0] = product2[i+1]
		} else if i == (n - 1) {
			nums[i] = product[i-1]
		} else {
			nums[i] = product[i-1] * product2[i+1]
		}
	}

	return nums
}



func productExceptSelf2(nums []int) []int {
	if len(nums) == 0 {
		return []int{}
	}

	if len(nums) == 1 {
		return []int{1}
	}
	n := len(nums)
	Left := make([]int, n)
	Right := make([]int, n)
	Left[0] = 1
	Right[n-1] = 1

	for i := 1; i < n; i++ {
		Left[i] = Left[i-1] * nums[i-1]
		Right[n-i-1] = Right[n-i] * nums[n-i]
	}

	for i := 0; i < n; i++ {
		nums[i] = Left[i] * Right[i]
	}

	return nums
}

func productExceptSelf3(nums []int) []int {
	if len(nums) == 0 {
		return []int{}
	}

	if len(nums) == 1 {
		return []int{1}
	}
	n := len(nums)
	Left := make([]int, n)
	Left[0] = 1
	for i := 1; i < n; i++ {
		Left[i] = Left[i-1] * nums[i-1]
	}

	fmt.Println("Left is ",Left)

	R := 1
	for i := n - 1; i >= 0; i-- {
		// 对于索引 i，左边的乘积为 answer[i]，右边的乘积为 R
		Left[i] = Left[i] * R
		// R 需要包含右边所有的乘积，所以计算下一个结果时需要将当前值乘到 R 上
		R *= nums[i]
	}


	return Left
}








func main() {
	//fmt.Println(productExceptSelf([]int{1, 2, 3, 4, 5}))
	//fmt.Println(productExceptSelf2([]int{1, 2, 3, 4, 5}))
	fmt.Println(productExceptSelf3([]int{1, 2, 3, 4}))

}
