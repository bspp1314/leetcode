package main

import "fmt"

type TestStruct struct{}

type TestInterFace interface {
	Hello()
	Hell2()
}


func NilOrNot(v interface{}) bool {
	return v == nil
}

func main() {
	var s *TestStruct
	var v TestInterFace
	fmt.Println(s == nil)
	fmt.Println(NilOrNot(s))
	fmt.Println(v == nil)
	fmt.Println(NilOrNot(v))
}




func rotate(nums []int, k int) {
	if len(nums) == 0 {
		return
	}

	k = k % len(nums)
	if k == 0 {
		return
	}
	// 1 2 3 4 / 5 6 7
	// 4 3 2 1 / 7 6 5
	// 5 6 7 1 2 3 4
	// right is 7 - 3 -1
	rotateHelp(nums, 0, len(nums)-k-1)
	rotateHelp(nums, len(nums)-k, len(nums)-1)
	rotateHelp(nums, 0, len(nums)-1)
}

func rotate2(nums []int, k int) {
	if len(nums) == 0 {
		return
	}

	l := len(nums)
	k = k % l
	if k == 0 {
		return
	}

	// k = 3
	// 1 2 3 4 5 6 7
	//swap index
	// 3 4 5 6 0 1 2
	// 1 2 3 4
	count := 0
	for i := 0; count < len(nums); i++ {
		currentIndex := i
		currentValue := nums[i]

		for {
			swapIndex := (currentIndex + k) % l
			temp := nums[swapIndex]
			nums[swapIndex] = currentValue
			currentIndex = swapIndex
			currentValue = temp
			count++
			//阻止其调回原来的位置
			if currentIndex == i {
				break
			}
		}
	}

}

func rotateHelp(nums []int, left, right int) {
	for left < right {
		nums[left], nums[right] = nums[right], nums[left]
		left = left + 1
		right = right - 1
	}
}

func main() {
	//nums1 := []int{-1, 100, 2, 99, 4, 5, 6, 7}
	//nums2 := []int{1, 2, 3, 4, 5, 6, 7}
	//nums3 := []int{1, 2, 3, 4, 5, 6, 7}
	//rotate2(nums1, 3)
	////rotate2(nums2,8)
	////rotate2(nums3,9)
	//rotate2(nil, 10)
	//fmt.Println(nums1)
	//fmt.Println(nums2)
	//fmt.Println(nums3)

	var s *TestStruct
	var v TestInterFace
	fmt.Println(s == nil)      // #=> true
	fmt.Println(NilOrNot(s))   // #=> false
	fmt.Println(v == nil)      // #=> true
	fmt.Println(NilOrNot(v))   // #=> false
}
