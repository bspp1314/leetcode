package main

import "fmt"
// 1 2 3 4 5 6 7 7 7 8 9
func searchRange(nums []int, target int) []int {
	res := []int{-1, -1}
	if len(nums) == 0 {
		return res
	}

	return search(nums, 0, len(nums)-1, target)
}

func searchFirst(nums []int,target int ) int  {
	left := 0
	right := len(nums) -1
	res := -1

	for left <= right {
		mid := left + (right-left) /2
		if nums[mid] == target {
			res = mid
			right = mid -1
		}else if nums[mid] > target {
			right = mid -1
		}else{
			left = mid +1
		}

	}

	return res
}

func searchLast(nums []int,target int ) int  {
	left := 0
	right := len(nums) -1
	res := -1

	for left <= right {
		mid := left + (right-left) /2
		if nums[mid] == target {
			res = mid
			left = mid + 1
		}else if nums[mid] > target {
			right = mid -1
		}else{
			left = mid +1
		}

	}

	return res
}


func searchRange2(nums []int, target int) []int {
	res := searchFirst(nums,target)
	if res == -1 {
		return []int{-1,-1}
	}
	res2 := searchLast(nums,target)
	return []int{res,res2}
}
func searchHelp(nums []int, left int, right int, taget int) []int {
	res := []int{-1, -1}
	for left <= right {
		mid := left + (right-left)

		if taget == nums[mid] {
			//get left range
			res[0] = mid
			res[1] = mid

			leftRange := search(nums, left, mid-1, taget)
			rightRange := search(nums, mid+1, right, taget)

			if leftRange[0] != -1 {
				res[0] = leftRange[0]
			}

			if rightRange[1] != -1 {
				res[1] = rightRange[1]
			}

			return res
		}else if taget > nums[mid] {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	return res
}


func search(nums []int, left int, right int, taget int) []int {
	res := []int{-1, -1}

	for left <= right {
		mid := left + (right-left)/2

		if taget == nums[mid] {
			res[0] = mid
			res[1] = mid
			leftRange := search(nums, left, mid-1, taget)
			rightRange := search(nums, mid+1, right, taget)

			if leftRange[0] != -1 {
				res[0] = leftRange[0]
			}

			if rightRange[1] != -1 {
				res[1] = rightRange[1]
			}

			return res
		} else if taget > nums[mid] {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	return res
}

func main() {
	fmt.Println("vim-go")
}
