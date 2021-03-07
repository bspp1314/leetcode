package main

import "fmt"

func containsNearbyDuplicate(nums []int, k int) bool {
	valueMap := make(map[int]int)

	for i := 0; i < len(nums); i++ {
		if _, ok := valueMap[nums[i]]; ok {
			return true
		}
		valueMap[nums[i]] = i

		if len(valueMap) > k {
			delete(valueMap, nums[i-k])
		}
	}

	return false
}

func containsNearbyDuplicate2(nums []int, k int) bool {
	duplicateMap := make(map[int]int,0)

	for i := 0; i < len(nums); i++ {
		v,exist := duplicateMap[nums[i]]
		if exist {
			if (i - v) <= k {
				return true
			}
		}

		duplicateMap[nums[i]] = i
	}

	return false
}

func abs(a int) int {
	if a < 0 {
		return -a
	}

	return a
}

func main() {
	fmt.Println(containsNearbyDuplicate([]int{1, 2, 3, 1, 2, 3}, 2))
	fmt.Println(containsNearbyDuplicate([]int{99, 99}, 2))
	fmt.Println(containsNearbyDuplicate([]int{1, 2, 3, 1}, 3))
}
