package main

import (
	"fmt"
	"sort"
)

// n >= 2
// 4 5 6 7 8   4个数之和
func mSumDfs(left int, nums []int, subList []int, m int, target int, res *[][]int) {
	if (m - len(subList)) == 2 {
		right := len(nums) - 1
		for left < right {
			sum := nums[left] + nums[right]
			if sum == target {
				temp := make([]int, len(subList), len(subList)+2)
				copy(temp, subList)
				temp = append(temp, nums[left])
				temp = append(temp, nums[right])
				*res = append(*res, temp)

				for left < right && nums[left] == nums[left+1] { // 去重
					left++
				}

				for left < right && nums[right] == nums[right-1] { // 去重
					right--
				}
				left++
				right--
			} else if sum > target {
				for left < right && nums[right] == nums[right-1] { // 去重
					right--
				}
				right--
			} else {
				for left < right && nums[left] == nums[left+1] { // 去重
					left++
				}
				left++
			}
		}
	}

	for i := left; i < len(nums)-(m-len(subList)-1); i++ {
		if nums[i] > target { // 剪枝操作，当当前数值大于目标值，则后续无需遍历
			return
		}

		if i > left && nums[i] == nums[i-1] {
			continue
		}

		zz := append(subList, nums[i])
		mSumDfs(i+1, nums, zz, m, target-nums[i], res)
	}
}

func mSum(nums []int, target int,m int) [][]int {
	res := make([][]int, 0)
	if len(nums) < m {
		return res
	}

	subList := make([]int, 0)
	sort.Sort(sort.IntSlice(nums))

	mSumDfs(0, nums, subList, m, target, &res)
	return res
}



func fourSum(nums []int, target int) [][]int {
	res := make([][]int, 0)
	if len(nums) < 2 {
		return res
	}

	//-1, 0, 1, 2, -1, -4
	//-4  -1 -1 0
	sort.Sort(sort.IntSlice(nums))

	for i := 0; i < len(nums)-3; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		for j := i + 1; j < len(nums)-2; j++ {
			if j > i+1 && nums[j] == nums[j-1] {
				continue
			}

			left := j + 1
			right := len(nums) - 1

			for left < right {
				sum := nums[i] + nums[j] + nums[left] + nums[right]
				if sum == 0 {
					res = append(res, []int{nums[i], nums[j], nums[left], nums[right]})

					for left < right && nums[left] == nums[left+1] { // 去重
						left++
					}

					for left < right && nums[right] == nums[right-1] { // 去重
						right--
					}
					left++
					right--
				} else if sum > 0 {
					for left < right && nums[right] == nums[right-1] { // 去重
						right--
					}
					right--
				} else {
					for left < right && nums[left] == nums[left+1] { // 去重
						left++
					}
					left++
				}
			}
		}
	}

	return res
}
func main() {
	out := mSum([]int{1, 0, -1, 0, -2, 2}, 0,4)
	fmt.Println(out)
}
