package main

import "fmt"

func permute(nums []int) [][]int {
	if len(nums) == 0 {
		return [][]int{}
	}

	if len(nums) == 1 {
		return [][]int{[]int{nums[0]}}
	}

	res := make([][]int, 0)
	for i := 0; i < len(nums); i++ {
		next := make([]int, len(nums)-1)
		copy(next, nums[:i])
		copy(next[i:], nums[i+1:])

		out := permute(nums)
		for _, v := range out {
			res = append(res, append(v, nums[i]))
		}
	}

	return res
}


func permute2(nums []int) [][]int {
	var result [][]int
	var visited = make([]bool, len(nums))
	var dfsHelp func(sub []int)

	dfsHelp = func(sub []int) {
		if len(nums) == len(sub) {
			tmp := make([]int, len(sub))
			copy(tmp, sub)
			result = append(result, tmp)
			return
		}

		for i := 0; i < len(nums); i++ {
			if !visited[i] {
				visited[i] = true

				newSub := append(sub, nums[i])
				dfsHelp(newSub)
				visited[i] = false
			}
		}
	}
	dfsHelp([]int{})
	return result
}

func main() {
	fmt.Println(permute2([]int{1, 2, 3}))
}
