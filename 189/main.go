package main

type TestStruct struct{}


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

func rotate2(nums []int, k int)  {
	l := len(nums)
	k  = k % l


	times := 0

	for i := 0;times < len(nums);i++ {
		index := i
		indexV := nums[i]


		for {
			index = (index + k ) % l
			tem  := nums[index]

			nums[index] = indexV
			indexV = tem
			times++
			if indexV == i {
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

}
