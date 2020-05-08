package main

import "fmt"

func merge(nums1 []int, m int, nums2 []int, n int) []int {
	//copy num1
	nums1Copy := make([]int, len(nums1))
	copy(nums1Copy, nums1)

	i := 0
	j := 0
	k := 0

	for i < m && j < n {
		if nums1Copy[i] < nums2[j] {
			nums1[k] = nums1Copy[i]
			i++
		} else {
			nums1[k] = nums2[j]
			j++
		}
		k++
	}

	for i < m {
		nums1[k] = nums1Copy[i]
		i++
		k++

	}

	for j < n {
		nums1[k] = nums2[j]
		j++
		k++
	}

	return nums1
}

func merge2(nums1 []int, m int, nums2 []int, n int) []int {
	p := len(nums1) - 1
	p1 := m - 1
	p2 := n - 1

	for p1 >= 0 && p2 >= 0 {
		if nums1[p1] > nums2[p2] {
			nums1[p] = nums1[p1]
			p1--
		} else {
			nums1[p] = nums2[p2]
			p2--
		}
		p--
	}

	for p1 >= 0 {
		nums1[p] = nums1[p1]
		p1--
		p--
	}

	for p2 >= 0 {
		nums1[p] = nums2[p2]
		p2--
		p--
	}
	return nums1
}

func main() {
	fmt.Println(merge2([]int{0}, 0, []int{1}, 1))
	fmt.Println("vim-go")
}
