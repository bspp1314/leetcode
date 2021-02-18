package main

import (
	"fmt"
	"sort"
)

type Intervals [][]int

func (s Intervals)Len() int  {
	return len(s)
}

func (s Intervals)Less(i,j int) bool  {
	return s[i][0] < s[j][0]
}

func (s Intervals)Swap(i,j int)  {
	s[i],s[j] = s[j],s[i]
}

func merge(intervals [][]int) [][]int {
	if len(intervals) <= 1 {
		return intervals
	}

	//sort
	sort.Sort(Intervals(intervals))

	res := [][]int{intervals[0]}
	for i := 1; i < len(intervals); i++ {
		if intervals[i][0] <= res[len(res)-1][1] {
			//merge
			if intervals[i][1] > res[len(res)-1][1] {
				res[len(res)-1][1] = intervals[i][1]
			}
		}else{
			res = append(res,intervals[i])
		}
	}

	return res
}

func Copy(i []int) []int  {
	res := make([]int,len(i))
	copy(res,i)
	return res
}

func main() {
	out := merge([][]int{
		{1,3},
		{2,6},
		{3,5},
		{8,10},
		{15,18},
	})

	fmt.Println(out)
}
