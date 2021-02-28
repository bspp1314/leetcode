package main

import (
	"fmt"
	"sort"
	"strconv"
)

type ans []string

func (n ans) Len() int {
	return len(n)
}

func (n ans) Swap(i, j int) {
	n[i], n[j] = n[j], n[i]
}

func (self ans) Less(i, j int) bool {
	a, _ := strconv.Atoi(self[i] + self[j])
	b, _ := strconv.Atoi(self[j] + self[i])
	return a > b
}

func largestNumber(nums []int) string {
	var a ans
	a = make([]string, len(nums))
	for i, v := range nums {
		a[i] = strconv.Itoa(v)
	}

	sort.Sort(a)
	res := ""
	if a[0] == "0" {
		return "0"
	}
	for _, v := range a {
		res += v
	}
	return res

}

func main() {
	//fmt.Println(largestNumber([]int{8, 7, 9}))
	fmt.Println(largestNumber([]int{3, 30, 34, 5, 9}))

}
