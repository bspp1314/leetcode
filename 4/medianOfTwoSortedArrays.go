package main

import (
	"fmt"
)

func findKthSmallest(num1 []int, num2 []int, k int) float64 {
	l1 := len(num1)
	l2 := len(num2)
	if l1 == 0 {
		return float64(num2[k-1])
	} else if l2 == 0 {
		return float64(num1[k-1])
	}

	if k == 1 {
		if num1[0] < num2[0] {
			return float64(num1[0])
		} else {
			return float64(num2[0])
		}
	}

	idx1 := (l1 / (l1 + l2)) * (k - 1)
	idx2 := (k - (idx1 + 1) - 1)

	if num1[idx1] == num2[idx2] {
		return float64(num1[idx1])
	} else if num1[idx1] < num2[idx2] {
		return findKthSmallest(num1[idx1+1:], num2, k-idx1-1)
	} else {
		return findKthSmallest(num1, num2[idx2+1:], k-idx2-1)
	}

}

func findMedianSortedArrays2(num1 []int, num2 []int) float64 {
	l1 := len(num1)
	l2 := len(num2)

	mid := (l1 + l2) / 2
	if ((l1 + l2) % 2) == 1 {
		return findKthSmallest(num1, num2, mid+1)
	} else {
		return (findKthSmallest(num1, num2, mid) + findKthSmallest(num1, num2, mid+1)) / 2.0
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
