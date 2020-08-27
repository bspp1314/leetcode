package main

import "sort"

func fourSum(nums []int, target int) [][]int {
	res := make([][]int, 0)
	if len(nums) < 2 {
		return res
	}

	//-1, 0, 1, 2, -1, -4
	//-4  -1 -1 0
	sort.Sort(sort.IntSlice(nums))

	for i:=0;i<len(nums)-3;i++{
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		for j:=i+1;j < len(nums)-2;j++ {
			if j > i && nums[j] == nums[j-1] {
				continue
			}

			left := j + 1
			right := len(nums) -1

			for left < right {
				sum := nums[i] + nums[j] + nums[left] + nums[right]
				if sum == 0 {
					res = append(res,[]int{nums[i],nums[j],nums[left],nums[right]})

					for left < right && nums[left] == nums[left+1] { // 去重
						left++
					}

					for left < right && nums[right] == nums[right-1] { // 去重
						right--
					}
					left++
					right--
				}else if sum > 0 {
					right--
				}else{
					left ++
				}
			}
		}
	}

	return res
}
func main()  {
	
}
