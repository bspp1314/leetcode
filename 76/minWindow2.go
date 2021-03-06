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
	cn := 0
	tMap := make([]int, 256)
	for i := 0; i < tLen; i++ {
		if tMap[t[i]] == 0 {
			cn++
		}
		tMap[t[i]]++
	}

	var res string
	windowCounts := make([]int, 256)
	for right < sLen {
		windowCounts[s[right]]++
		v := windowCounts[s[right]]
		if v == tMap[s[right]] {
			formed++
		}

		if formed == cn {
			newLen := right - left + 1
			if newLen > len(res) {
				res = s[left : right+1]
			}

			if tMap[s[left]] > 0 {
				if tMap[s[left]] == windowCounts[s[left]] {
					formed--
				}
				windowCounts[s[left]]--
			}
			left++
		}

		right++
	}
	return res
}
func main() {
	//e := "AAAA"
	fmt.Println(minWindow("ADOBECODEBANC", "ABC"))
	fmt.Println(minWindow("A", "AA"))
	fmt.Println(minWindow("A", "A"))
	fmt.Println(minWindow("aaaaaaaaaaabbbbbcdd", "abcdd"))
	//fmt.Println("vim-go")
}
