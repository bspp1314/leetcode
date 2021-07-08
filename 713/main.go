package main

func numSubarrayProductLessThanK(nums []int, k int) int {
	if len(nums) == 0 {
		return 0
	}

	if k <= 1 {
		return 0
	}

	left := 0
	product := 1
	res := 0

	for right := 0; right < len(nums); right++ {
		product = product * nums[right]
		for product >= k {
			product = product / nums[left]
			left++
		}

		res += right - left + 1
	}

	return res
}


// 5 2 3 4 5
// 5
// 52 2
// 523 23 3
// 5234 2 3 4
// 2345 345 35 5

func main() {
	numSubarrayProductLessThanK([]int{1,2,3}, 1)
}
