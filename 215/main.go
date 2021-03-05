package main

import (
	"fmt"
	"math/rand"
	"time"
)

func findKthLargest(nums []int, k int) int {
	rand.Seed(time.Now().UnixNano())
	return sort(nums, 0, len(nums)-1, k-1)
}

func sort(nums []int, left int, right int,k int) int  {
	index := patition(nums,left,right)
	if index == k   {
		return nums[k]
	}else if  index < k {
		return sort(nums,index + 1,right,k)
	}else{
		return sort(nums,left,index-1,k)
	}
}
func randomPartition(l, r int) int {
	return rand.Int()%(r-l+1) + l
}

func patition(nums []int, l int, r int) int {
	index := randomPartition(l, r)
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
	a := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10,10}

	fmt.Println(findKthLargest(a,3))
}
