package main

import (
	"fmt"
	"sort"
	"sync"
)

type Envelopes [][]int

func (e Envelopes) Len() int {
	return len(e)
}

func (e Envelopes) Swap(i, j int) {
	e[i], e[j] = e[j], e[i]
}

func (e Envelopes) Less(i, j int) bool {
	if e[i][0] == e[j][0] {
		return e[j][1] < e[i][1]
	}else{
		return e[i][0] < e[j][0]
	}
}

func lengthOfLIS(arr []int) int {
	var result []int
	for i := range arr {
		if len(result) == 0 || arr[i] > result[len(result) - 1 ]{
			result = append(result, arr[i])
		}else {
			left, right := 0, len(result) - 1
			for left < right {
				mid := left + (right-left)/2
				if result[mid] < arr[i] {
					left = mid+1
				}else {
					right = mid
				}
			}
			result[left] = arr[i]
		}
	}
	return len(result)
}


func maxEnvelopes(envelopes [][]int) int {
	if len(envelopes) <= 1 {
		return len(envelopes)
	}

	sort.Sort(Envelopes(envelopes))

	nums := make([]int,len(envelopes))
	for i := 0; i < len(envelopes); i++ {
		nums[i] = envelopes[i][1]
	}


	return lengthOfLIS(nums)
}


func main() {
	//out := maxEnvelopes([][]int{[]int{30, 50}, []int{12, 2}, []int{3, 4}, []int{12, 15}})
	//log.Println(out)


	for i := 0;i<10000;i++ {
		var sw sync.WaitGroup
		var x,y int
		sw.Add(2)
		go func() {
			defer sw.Done()
			x = 1
			fmt.Print(y)
		}()

		go func() {
			defer sw.Done()
			y = 1
			fmt.Print(x)
		}()

		fmt.Println()

		sw.Wait()
	}
}


