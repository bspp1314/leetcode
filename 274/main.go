package main

import (
	"fmt"
	"sort"
)

func Min(a,b int) int   {
	if a < b {
		return a
	}

	return b
}

func hIndex(citations []int) int {
	sort.Slice(citations, func(i, j int) bool {
		return citations[i] < citations[j]
	})


	for i := 0; i < len(citations); i++ {
		h := len(citations) - i
		if h <= citations[i]  {
			return h
		}
	}

	return 0
}

func hIndex2(citations []int) int   {
	left := 0
	right := len(citations) - 1

	for left <= right {
		mid := left + (right- left ) / 2

		h := len(citations) - mid
		if citations[mid] >= h {
			right = mid
		}else{
			left = mid +1
		}
	}

	return len(citations) - left
}
func main()  {
 	fmt.Println(hIndex([]int{0}))
	fmt.Println(hIndex([]int{3,3,3,4,4,5,100,101,101}))
}
