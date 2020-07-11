package main

import "fmt"

//假设按照升序排序的数组在预先未知的某个点上进行了旋转。
//
//( 例如，数组 [0,1,2,4,5,6,7] 可能变为 [4,5,6,7,0,1,2] )。
//
//搜索一个给定的目标值，如果数组中存在这个目标值，则返回它的索引，否则返回 -1 。
//
//你可以假设数组中存在重复的元素。
//
//你的算法时间复杂度必须是 O(log n) 级别。
//
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
func search(nums []int, target int) bool {
	if len(nums) == 0 {
		return false
	}

	if len(nums) == 1 {
		return nums[0] == target
	}

	left := 0
	right := len(nums) - 1

	for left <= right {
		for left < right && nums[left] == nums[left+1] {
			left = left + 1
		}

		for right > left && nums[right] == nums[right-1] {
			right = right - 1
		}
		mid := (left + right) / 2
		if nums[mid] == target {
			return true
		}

		// mid 的值在左边
		if nums[left] <= nums[mid] {
			// nums[0] <= taget
			if nums[left] <= target && target < nums[mid] {
				right = mid - 1
			} else {
				left = mid + 1
			}
		} else {
			if nums[mid] < target && target <= nums[right] {
				left = mid + 1
			} else {
				right = mid - 1
			}
		}
	}

	return false
}

func main() {
	nums := []int{5, 5, 6, 6, 7, 7, 8, 8, 9, 9, 1, 1, 2, 2, 3, 3, 4, 4}
	for i := 0; i < len(nums); i++ {
		fmt.Printf(" %d  %v ", nums[i], search(nums, nums[i]))
	}
	fmt.Println()

	for i := 0; i < len(nums); i++ {
		fmt.Printf(" %v ", search(nums, nums[i]+9))
	}
	fmt.Println()

	for i := 0; i < len(nums); i++ {
		fmt.Printf(" %v ", search(nums, nums[i]-9))
	}
	fmt.Println()
}
