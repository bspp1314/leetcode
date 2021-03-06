package main

import (
	"log"
)

//左右指针滑动窗口
func lengthOfLongestSubstring(s string) int {
	sLen := len(s)
	l := 0
	right := 0
	left := 0
	//保存某一字符最新出现的位置
	array := make([]int, 256)

	for right < sLen {
		//出现重复字符，左指针移动
		if array[s[right]] > left {
			left = array[s[right]]
		}
		if l < right-left+1 {
			l = right - left + 1
		}
		array[s[right]] = right + 1
		right++
	}
	return l
}

func Max(a,b int) int {
	if a > b {
		return a
	}

	return b
}

func lengthOfLongestSubstring2(s string) int {
	if len(s) <= 1 {
		return len(s)
	}

	left := 0
	right := 0
	n := len(s)
	max := 0
	kMap := make(map[byte]int)

	for right < n {
		//exit
		index  := kMap[s[right]]
		if  index > left {
			left = index
		}
		max = Max(right-left+1,max)
		kMap[s[right]] = right+1
		right++
	}

	return max
}

func main() {
	log.Println(lengthOfLongestSubstring("abba"))
	log.Println(lengthOfLongestSubstring2("abba"))
}
