package main

import "fmt"

func max(i, j int) int {
	if i > j {
		return i
	}

	return j
}
func canJump1(nums []int) bool {
	n := len(nums)
	rightmost := 0
	for i := 0; i < n; i++ {
		if i <= rightmost {
			rightmost = max(rightmost, i+nums[i])
			if rightmost >= n-1 {
				return true
			}
		}
	}

	return false

}

func canJump2(nums []int) bool{
	if len(nums) == 0 {
		return true
	}

	jumpMap := make(map[int]bool)
	lastIndex := len(nums) - 1

	var dfs func(begin int) bool

	dfs = func(begin int) (res bool) {
		if begin >= lastIndex {
			jumpMap[lastIndex] = true
			return true
		}

		if v, ok := jumpMap[begin]; ok {
			return v
		}

		if nums[begin] == 0 {
			jumpMap[begin] = false
			return false
		}

		for j := nums[begin];j >=1; j-- {
			v := dfs(begin + j)
			if v {
				jumpMap[begin] = v
				return v
			}
		}

		return false
	}

	return dfs(0)
}

func main() {
	fmt.Println(canJump2([]int{2,3,1,1,4}))
}
