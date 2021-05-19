package main

import (
	"fmt"
	"math"
)

func findK(nums1 []int,nums2 []int,i,j,k int) float64  {
	if i >= len(nums1) {
		return float64(nums2[j+k-1])
	}

	if j >= len(nums2) {
		return float64(nums2[i+k-1])
	}

	if k == 1 {
		return math.Min(float64(nums1[i]),float64(nums2[j]))
	}

	mid1 := 0
	mid2 := 0

	if i + k / 2  -1  < len(nums1) {
		mid1 = nums1[i+k/2-1]
	}else{
		mid1 = math.MaxInt64
	}

	//计算出每次要比较的两个数的值，来决定 "删除"" 哪边的元素
	if j + k/2 -1 < len(nums2) {
		mid2 = nums2[j+k/2-1]
	}else{
		mid2 = math.MaxInt64
	}

	//通过递归的方式，来模拟删除掉前K/2个元素
	if mid1 < mid2 {
		return findK(nums1,nums2,i+k/2,j,k-k/2)
	}else{
		return findK(nums1,nums2,i,j+k/2,k-k/2)
	}
}

func findMedianSortedArrays2(num1 []int, num2 []int) float64 {
	l1 := len(num1)
	l2 := len(num2)

	mid := (l1 + l2) / 2
	if ((l1 + l2) % 2) == 1 {
		return findK(num1, num2,0,0, mid+1)
	} else {
		return (findK(num1, num2,0,0, mid+1) + findK(num1, num2, 0,0,mid+1)) / 2.0
	}
}

func findMedianSortedArrays(num1 []int, num2 []int) float64 {
	len1 := len(num1)
	len2 := len(num2)
	i := 0
	j := 0
	k := 0

	num3 := make([]int, (len1 + len2))
	for i < len1 && j < len2 {
		if num1[i] < num2[j] {
			num3[k] = num1[i]
			i++
		} else {
			num3[k] = num2[j]
			j++
		}
		k++
	}
	for i < len1 {
		num3[k] = num1[i]
		i++
		k++
	}
	for j < len2 {
		num3[k] = num2[j]
		j++
		k++
	}

	if (len1+len2)%2 == 1 {
		return float64(num3[(len1+len2)/2])
	} else {
		v1 := float64(num3[(len1+len2)/2-1])
		v2 := float64(num3[(len1+len2)/2])
		return (v1 + v2) / 2.0
	}

}

func main() {
	num1 := []int{1}
	num2 := []int{2}
	fmt.Println(findMedianSortedArrays(num1, num2))
	fmt.Println(findMedianSortedArrays2(num1, num2))
}
