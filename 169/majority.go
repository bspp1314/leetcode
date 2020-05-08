package main

import "fmt"

func majorityElement(nums []int) int {
	max_time := 0
	max_value := 0
	timesMap := make(map[int]int)

	n1 := len(nums) / 2

	for i := 0; i < len(nums); i++ {
		fmt.Println("i", i)
		if timesMap[nums[i]] != 0 {
			timesMap[nums[i]] = timesMap[nums[i]] + 1
		} else {
			timesMap[nums[i]] = 1
		}

		if timesMap[nums[i]] > max_time {
			max_time = timesMap[nums[i]]
			max_value = nums[i]
		}

		if timesMap[nums[i]] > n1 {
			return nums[i]
		}

	}

	fmt.Println(timesMap)

	return max_value
}

func DividemajorityElement(nums []int) int {
	return majorityElementRec(nums, 0, len(nums)-1)
}

func countInRange(nums []int, num, lo, hi int) int {
	res := 0
	for i := lo; i <= hi; i++ {
		if nums[i] == num {
			res++
		}
	}

	return res
}

func majorityElementRec(nums []int, lo, hi int) int {
	if lo == hi {
		return nums[lo]
	}

	mid := (hi-lo)/2 + lo

	left := majorityElementRec(nums, lo, mid)
	right := majorityElementRec(nums, mid+1, hi)

	if left == right {
		return left
	}

	leftCount := countInRange(nums, left, lo, hi)
	rightCount := countInRange(nums, right, lo, hi)

	if leftCount > right {
		return leftCount
	} else {
		return rightCount
	}
}

func main() {
	//fmt.Println(majorityElement([]int{2,2,1,1,1,2,2}))
	fmt.Println(DividemajorityElement([]int{2, 2, 1, 1, 1, 2, 2}))

}
