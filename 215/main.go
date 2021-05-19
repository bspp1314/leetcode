package main

import (
	"math/rand"
	"time"
)

func findKthLargest(nums []int, k int) int {
	rand.Seed(time.Now().UnixNano())
	return sort(nums, 0, len(nums)-1, k-1)
}

func sort(nums []int, left int, right int, k int) int {
	index := patition(nums, left, right)
	if index == k {
		return nums[k]
	} else if index < k {
		return sort(nums, index+1, right, k)
	} else {
		return sort(nums, left, index-1, k)
	}
}

func findKthLargestZ(nums []int, k int) []int {
	Sort2(nums, 0, len(nums)-1, k-1)
	return nums[:k]

}

func Sort2(nums []int, left int, right int, k int) {
	if left >= right {
		return
	}
	index := patition(nums, left, right)
	if index >= k {
		Sort2(nums, left, index-1, k)
	} else {
		Sort2(nums, left, index-1, k)
		Sort2(nums, index+1, right, k)
	}
}

func randomPartition(l, r int) int {
	return rand.Int()%(r-l+1) + l
}

func patition(nums []int, l int, r int) int {
	index := randomPartition(l, r)
	index = l
	p := nums[index]
	//为了获取中间的数
	nums[index], nums[l] = nums[l], nums[index]
	i := l
	j := r

	for i < j {
		for nums[j] <= p && i < j {
			j--
		}

		for nums[i] >= p && i < j {
			i++
		}

		nums[i], nums[j] = nums[j], nums[i]
	}

	nums[l], nums[i] = nums[i], nums[l]

	return i
}

func main() {
}
