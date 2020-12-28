package main

import (
	"fmt"
)

type Degree struct {
	Begin int
	End   int
	Degree int
}

func findShortestSubArray(nums []int) int {

	if len(nums) == 0 {
		return 0
	}
	degree := make(map[int]*Degree)
	maxDegree := 0
	res := 0
	for i:=0;i<len(nums);i++ {
		d,ok := degree[nums[i]]
		if !ok {
			d = &Degree{}
		}

		if d.Degree == 0  {
			d.Begin = i
		}

		d.End = i
		d.Degree++

		if  d.Degree > maxDegree {
			maxDegree = d.Degree
			res = d.End - d.Begin + 1
		}else if d.Degree == maxDegree {
			newRes := d.End - d.Begin + 1
			if newRes < res {
				res = newRes
			}
		}
		degree[nums[i]] = d
	}

	return res
}

func main() {
	fmt.Print(findShortestSubArray([]int{1,2,2,1,2,1}))
}
