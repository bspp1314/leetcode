package main

import "fmt"

func maxSubArray(nums []int) int {
	if len(nums) <= 0 {
		return 0
	}

	max := nums[0]
	sum := nums[0]

	for i:=1;i <len(nums);i++ {
		sum += nums[i]
		if sum < nums[i] {
			sum = nums[i]
		}

		if max < sum {
			max = sum
		}
	}

	return max
}

func maxSubArray3(nums []int) int {
	if len(nums) <= 0 {
		return 0
	}

	max := nums[0]
	sum := nums[0]
	i := 1
	l := len(nums)

	for i < l  {
		sum += nums[i]
		if sum < nums[i] {
			sum = nums[i]
		}

		if max < sum {
			max = sum
		}
		i++
	}

	return max
}

func maxSubArray1(nums []int) int {
	if len(nums) <= 0 {
		return 0
	}

	if len(nums) == 1 {
		return nums[0]
	}

	dp := make([]int,len(nums))
	dp[0] = nums[0]

	//max := nums[0]
	//sum := nums[0]
	max := dp[0]

	for i:=1;i <len(nums);i++ {
		if dp[i-1] + nums[i]  > nums[i] {
			dp[i] = dp[i-1] + nums[i]
		}else{
			dp[i] =  nums[i]
		}

		if dp[i] > max {
			max = dp[i]
		}
	}


	return max
}


func maxSubArray2(nums []int) int {
	return get(nums, 0, len(nums) - 1).mSum;
}

func pushUp(l, r Status) Status {
	iSum := l.iSum + r.iSum
	lSum := max(l.lSum, l.iSum + r.lSum)
	rSum := max(r.rSum, r.iSum + l.rSum)
	mSum := max(max(l.mSum, r.mSum), l.rSum + r.lSum)
	return Status{lSum, rSum, mSum, iSum}
}

func get(nums []int, l, r int) Status {
	if (l == r) {
		return Status{nums[l], nums[l], nums[l], nums[l]}
	}
	m := (l + r) >> 1
	lSub := get(nums, l, m)
	rSub := get(nums, m + 1, r)
	return pushUp(lSub, rSub)
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

type Status struct {
	lSum, rSum, mSum, iSum int
}


func main() {
	//arr := []int{-2,1,-3,4,-1,2,1,-5,4}
	//s := get(arr,0,len(arr)-1)
	//fmt.Println(s)
	fmt.Println(maxSubArray([]int{-2,1,-3,4,-1,2,1,-5,4}))
	fmt.Println(maxSubArray1([]int{-2,1,-3,4,-1,2,1,-5,4}))
}
