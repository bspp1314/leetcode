package main

func removeDuplicates(nums []int) int {
	if len(nums) <= 2 {
		return len(nums)
	}
	j := 1
	count := 1

	for i := 1; i < len(nums); i++ {
		if nums[i] == nums[i-1] {
			count++
		}else{
			count = 1
		}

		if count <= 2 {
			nums[j] = nums[i]
			j += 1
		}
	}

	return 0
}
