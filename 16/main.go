package main

import (
	"fmt"
	"math"
	"sort"
)

func threeSumClosest(nums []int, target int) int {

	res := 0
	dis := math.MaxInt64 //距离
	abs := func(a int) int {
		if a < 0 {
			return -a
		}

		return a
	}

	sort.Sort(sort.IntSlice(nums))

	for i := 0; i < len(nums)-2; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		left := i + 1
		right := len(nums) - 1
		for left < right {
			sum := nums[left] + nums[right] + nums[i]
			newDis := abs(sum - target)

			if dis > newDis {
				dis = newDis
				res = sum
			}

			if dis == 0 {
				return res
			}

			if sum < target {
				for left < right && nums[left] == nums[left+1] { // 去重
					left++
				}
				left++
			} else {
				for left < right && nums[right] == nums[right-1] { // 去重
					right--
				}
				right--
			}
		}
	}

	return res
}

func main() {
	out := threeSumClosest([]int{0, 2, 1, -3}, 1)
	fmt.Println(out)
}
