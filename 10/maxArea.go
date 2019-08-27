package main

import (
	"fmt"
)

//给定 n 个非负整数 a1，a2，...，an，每个数代表坐标中的一个点 (i, ai) 。在坐标内画 n 条垂直线，垂直线 i 的两个端点分别为 (i, ai) 和 (i, 0)。找出其中的两条线，使得它们与 x 轴共同构成的容器可以容纳最多的水。

func max(i, j int) int {
	if i > j {
		return i
	} else {
		return j
	}
}
func min(i, j int) int {
	if i > j {
		return j
	} else {
		return i
	}
}

//暴力法
func maxArea(height []int) int {
	l := len(height)
	maxArea := 0

	for i := 0; i < l; i++ {
		for j := i + 1; j < l; j++ {
			newArea := (j - i) * min(height[i], height[j])
			maxArea = max(maxArea, newArea)
		}
	}
	return maxArea
}

// 双指针法
func maxArea2(height []int) int {
	i := 0
	j := len(height) - 1
	maxArea := 0
	for i < j {
		newArea := (j - i) * min(height[i], height[j])
		maxArea = max(maxArea, newArea)
		if height[i] < height[j] {
			i++
		} else {
			j--
		}
	}
	return maxArea
}
func main() {
	height := []int{1, 8, 6, 2, 5, 4, 8, 3, 7}
	fmt.Println(maxArea(height))
	fmt.Println(maxArea2(height))
}
