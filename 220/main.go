package main

import (
	"fmt"
)

// 3/10 =============> 0
// -3/10 need index is -1
func GetBucketIndex(x int, w int) int  {
		if x < 0 {
			return (x+1) /w -1
		}

		return x / w
}

func abs(a int ) int   {
	if a > 0 {
		return a
	}

	return -a
}

func containsNearbyAlmostDuplicate(nums []int, k int, t int) bool {
	if t < 0 {
		return false
	}

	bucket := make(map[int]int)
	w := t +1

	for i:=0;i<len(nums);i++ {
		//获取所在的桶位置
		m := GetBucketIndex(nums[i], w)

		_,ok := bucket[m]
		if ok {
			return true
		}

		v1,ok := bucket[m+1]
		if ok && abs(v1-nums[i]) < w {
			return true
		}


		v2,ok := bucket[m-1]
		if ok && abs(v2-nums[i]) < w {
			return true
		}

		bucket[m] = nums[i]

		if i >= k {
			delete(bucket,GetBucketIndex(nums[i-k], w))
		}
	}

	return false
}

func main() {

	fmt.Println("a ^ a ",2 ^ 4 ^ 4 )
	//fmt.Println(containsNearbyAlmostDuplicate([]int{1,5,9,1,5,9},2,3))
}
