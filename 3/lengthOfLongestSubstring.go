package main

import (
	"fmt"
	"log"
)

func lengthOfLongestSubstring(s string) int {
	size := len(s)
	max_len := 0
	begin := 0
	end := 0
	array := make([]int, 256)

	for i := 0; i < size; i++ {
		end = i
		//出现重复数据
		if array[s[i]] > begin {
			begin = array[s[i]]
		}

		if max_len < (end - begin + 1) {
			max_len = end - begin + 1
		}
		array[s[i]] = i + 1
	}
	return max_len
}

func main() {
	log.Println(lengthOfLongestSubstring("dddddddddddddddd888888888888d"))
	fmt.Println("vim-go")
}
