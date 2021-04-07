package main

import "fmt"

// 第一类

//这是最简单的一类，也是我们最开始学二分查找法需要解决的问题，比如我们有数组 [2, 4, 5, 6, 9]，target = 6，那么我们可以写出二分查找法的代码如下：

func find(a []int,target int) int   {
	left := 0
	right := len(a) -1
	for left < right {
		mid := left + (right - left) /2
		if a[mid] == target {
			return  mid
		}else if a[mid] < target {
			left = mid + 1
		}else{
			right = right -1
		}
	}

	return -1
}

// 查找第一个大于等于目标值的数，可变形为查找最后一个小于目标值的数
// 1 2 2 3 3 4 4  target 2  want index is 2
// 这一类可以轻松的变形为查找最后一个小于目标值的数，怎么变呢。
//我们已经找到了第一个不小于目标值的数，那么再往前退一位，返回 right - 1，就是最后一个小于目标值的数。

func find2(a []int,target int) int   {
	left := 0
	right := len(a) -1
	for left < right {
		mid := left + (right - left) /2
		if a[mid] < target {
			left = mid + 1
		}else{
			right = mid
		}
	}

	return right
}

// 第三类： 查找第一个大于目标值的数，可变形为查找最后一个不大于目标值的数
// 这一类可以轻松的变形为查找最后一个不大于目标值的数，怎么变呢。我们已经找到了第一个大于目标值的数，那么再往前退一位，返回 right - 1，
func find3(nums []int,target int) int   {
	left := 0
	right := len(nums) -1
	for left < right {
		mid := left + (right - left) / 2
		if nums[mid] <= target {
			left = mid + 1
		}else{
			right = mid
		}
	}

	return right
}

func main() {
	fmt.Println(find2([]int{1,2,2,3,3,4,4},5))
	fmt.Println(find3([]int{1,2,2,3,3,4,4},2))
}
