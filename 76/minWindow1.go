package main

import "fmt"

func minWindow(s string, t string) string {
	tLen := len(t)
	sLen := len(s)
	if sLen == 0 || sLen == 0 {
		return ""
	}

	formed := 0
	left := 0
	right := 0
	begin := 0
	l := 0
	tMap := make(map[byte]int, 0)
	for i := 0; i < tLen; i++ {
		tMap[t[i]]++
	}

	windowCounts := make(map[byte]int, 0)
	for right < sLen {
		windowCounts[s[right]]++
		if v := windowCounts[s[right]]; v == tMap[s[right]] {
			formed++
		}
		for left <= right && formed == len(tMap) {
			if l == 0 || l > right-left+1 {
				begin = left
				l = right - left + 1
			}
			if v1, ok := tMap[s[left]]; !ok {
				left++
			} else {
				if v1 == windowCounts[s[left]] {
					formed--
				}
				windowCounts[s[left]]--
				left++
			}
		}
		right++
	}
	return s[begin : begin+l]
}
func main() {
	//e := "AAAA"
	fmt.Println("0")
	fmt.Println(minWindow("ADOBECODEBANC", "ABC"))
	fmt.Println(minWindow("A", "AA"))
	fmt.Println(minWindow("A", "A"))
	fmt.Println(minWindow("aaaaaaaaaaabbbbbcdd", "abcdd"))
	//fmt.Println("vim-go")
}
