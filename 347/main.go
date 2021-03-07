package main

import (
	"fmt"
	"math/rand"
	"time"
)



func topKFrequent(nums []int, k int) []int {
	rand.Seed(time.Now().UnixNano())
	indexMap := make(map[int]int)
	values := make([][2]int, 0)

	for i := 0; i < len(nums); i++ {
		index, exist := indexMap[nums[i]]
		if !exist {
			values = append(values, [2]int{nums[i], 0})
			index = len(values) -1
			indexMap[nums[i]] = index
		}
		values[index][1] = values[index][1] + 1
	}



	sort(values,0,len(values)-1,k-1)

	var res []int
	for i := 0; i < k; i++ {
		res = append(res,values[i][0])
	}
	return res
}

func sort(nums [][2]int, left int, right int, k int)  {
	if left >= right {
		return
	}

	index := patition(nums,left,right)
	if index >= k {
		sort(nums,left,index-1,k)
	}else{
		sort(nums,left,index-1,k)
		sort(nums,index+1,right,k)
	}
}

func randomPartition(l, r int) int {
	return rand.Int()%(r-l+1) + l
}

func patition(nums [][2]int, l int, r int) int {
	index := randomPartition(l, r)
	p := nums[index][1]
	////为了获取中间的数
	nums[index], nums[l] = nums[l], nums[index]
	i := l
	j := r

	for i < j {
		for nums[j][1] <= p && i < j {
			j--
		}

		for nums[i][1] >= p && i < j {
			i++
		}

		nums[i], nums[j] = nums[j], nums[i]
	}

	nums[l], nums[i] = nums[i], nums[l]

	return i
}

func main() {
	fmt.Println(topKFrequent([]int{4,1,-1,2,-1,2,3},2))
}
