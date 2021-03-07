package main

import "fmt"

func summaryRanges(nums []int) []string {
	if len(nums) == 0 {
		return []string{}
	}

	if len(nums) == 1 {
		return []string{fmt.Sprintf("%d", nums[0])}
	}

	res := []string{}
	v := fmt.Sprintf("%d", nums[0])
	i:= 1
	for i = 1; i < len(nums); i++ {
		if nums[i] == (nums[i-1] + 1) {
			continue
		} else {
			if fmt.Sprintf("%d", nums[i-1]) != v {
				res = append(res, fmt.Sprintf("%s=>%d", v, nums[i-1]))
			} else {
				res = append(res, v)

			}
			v = fmt.Sprintf("%d", nums[i])
		}
	}

	if v != "" {
		if fmt.Sprintf("%d", nums[i-1]) != v {
			res = append(res, fmt.Sprintf("%s=>%d", v, nums[i-1]))
		} else {
			res = append(res, v)

		}
	}

	return res

}

func main() {
	fmt.Println(summaryRanges([]int{1,2,3,4,5,7,9,10,11,12,13}))
}
