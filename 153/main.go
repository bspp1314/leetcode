package main

import "fmt"

func findMin(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	if len(nums) == 1 {
		return nums[0]
	}

	left := 0
	right := len(nums) - 1

	for left < right {
		mid := left + (right-left)/2
		//这里有个编程技巧
		//因为l<r 所以最后一轮肯定是(r,r+1)
		//那么mid 肯定是取值l 当判断条件是mid与l比时 会出现与自身比 造成出现等于情况 不好判断
		//所以判断条件时mid 与 r比 这样肯定是不同的两个数比
		if nums[mid] < nums[right] { // mid可能为最小值
			right = mid
		} else if nums[mid] > nums[right]{
			left = mid + 1 // mid肯定不是最小值
		}else{
			right--
		}

	}

	return nums[left]
}
func main() {
	fmt.Println(findMin([]int{1, 1,2, 3, 4, 5,1,1,1}))
	//fmt.Println(findMin([]int{2, 3, 4, 5, 1}))
	//fmt.Println(findMin([]int{3, 4, 5, 1, 2}))
	//fmt.Println(findMin([]int{4, 5, 1, 2, 3}))
	//fmt.Println(findMin([]int{5, 1, 2, 3, 4}))
	//fmt.Println(findMin([]int{1, 5}))
	//fmt.Println(findMin([]int{5, 1}))

}
