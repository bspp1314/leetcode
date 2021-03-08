package main

import "fmt"


func maxSlidingWindow(nums []int, k int) []int {
	q := []int{}
	var push func(i int)
	push = func(i int) {
		if len(q) > 0 && nums[i] >= nums[q[len(q)-1]] {
			q = q[:len(q)-1]
			push(i)
		}else{
			q = append(q, i)
		}
	}

	for i := 0; i < k; i++ {
		push(i)
	}

	n := len(nums)
	ans := make([]int, 1, n-k+1)
	ans[0] = nums[q[0]]


	for i:=k;i < n;i++ {
		push(i)

		if q[0] <= i - k {
			q = q[1:]
		}

		ans = append(ans, nums[q[0]])

	}

	return ans

}

func main() {
	fmt.Println(maxSlidingWindow([]int{7,2,4},2))
}
