package main

import (
	"fmt"
	"math"
	"sort"
)

type Pairs [][]int

func (p Pairs) Len() int   {
	return len(p)
}

func (p Pairs) Less(i,j int) bool  {

	return p[i][0] < p[j][0]
}

func (p Pairs)Swap(i,j int)   {
	p[i],p[j] = p[j],p[i]
}

type Pairs2 [][]int

func (p Pairs2) Len() int   {
	return len(p)
}

func (p Pairs2) Less(i,j int) bool  {

	return p[i][0] < p[j][0]
}

func (p Pairs2)Swap(i,j int)   {
	p[i],p[j] = p[j],p[i]
}


func findLongestChain(pairs [][]int) int {
	if len(pairs) <= 1 {
		return len(pairs)
	}
	sort.Sort(Pairs(pairs))


	max := func(a, b int) int {
		if a >  b {
			return a
		}

		return b
	}

	// begin end
	// 义 dp[i] 为为考虑前 i个元素，以第 i 个数字结尾的最长数对链。
	dp := make([]int, len(pairs))


	dp[0] = 1
	res := 1
	for i := 1; i < len(pairs); i++ {
		dp[i] = 1
		for j := 0; j < i; j++ {
			if pairs[i][0] > pairs[j][1] {
				dp[i] = max(dp[i], dp[j]+1)
			}

		}

		res = max(res, dp[i])
	}

	return res
}

func findLongestChain2(pairs [][]int) int {
	if len(pairs) <= 1 {
		return len(pairs)
	}
	sort.Sort(Pairs(pairs))



	tailIndex := [][]int{pairs[0]}


	for i:=1;i<len(pairs);i++ {
		if pairs[i][0] > tailIndex[len(tailIndex)-1][1] {
			tailIndex = append(tailIndex,pairs[i])
		}else{
			left := 0
			right := len(tailIndex) -1

			for left < right {
				mid := left + (right - left)/2
				if pairs[i][1] > tailIndex[mid][1] {
					left = mid + 1
				} else if pairs[i][1] < tailIndex[mid][1] {
					right = mid
				}else{
					right = mid
					break
				}
			}


			if tailIndex[right][1] > pairs[i][1] {
				tailIndex[right] = pairs[i]
			}

		}
	}

	return len(tailIndex)
}


func findLongestChain3(pairs [][]int) int {
	sort.Sort(Pairs2(pairs))

	cur := math.MinInt64
	ans := 0

	for i := 1; i < len(pairs); i++ {
		if cur < pairs[i][0] {
			cur = pairs[i][1]
			ans++
		}
	}

	return  ans
}


func main() {
	//fmt.Println(findLongestChain([][]int{{1,2},{5,6},{7,8},{3,10}}))
	fmt.Println(findLongestChain2([][]int{{1,2},{5,6},{7,8},{3,10}}))
	fmt.Println(findLongestChain2([][]int{{3,4},{2,3},{1,2}}))

}
