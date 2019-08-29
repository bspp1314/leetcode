package main

import "fmt"

func findAnagrams(s string, p string) []int {
	sLen := len(s)
	pLen := len(p)
	if sLen == 0 || pLen == 0 {
		return nil
	}
	needs := make(map[byte]int)
	for i := 0; i < pLen; i++ {
		needs[p[i]]++
	}
	windows := make(map[byte]int)
	right := 0
	left := 0
	match := 0
	res := make([]int, 0)
	for right < sLen {
		c := s[right]
		if needs[c] > 0 {
			windows[c]++
			if windows[c] == needs[c] {
				match++
			}
		}
		right++
		for match == len(needs) {
			if right-left == pLen {
				res = append(res, left)
			}
			c2 := s[left]
			if needs[c2] > 0 {
				windows[c2]--
				if windows[c2] < needs[c2] {
					match--
				}
			}
			left++
		}

	}
	return res
}

func main() {
	//fmt.Println(findAnagrams("cbac", "bca"))
	//fmt.Println(findAnagrams("aabcaa", "abac"))
	//fmt.Println(findAnagrams("cbaebabacd", "bca"))
	//fmt.Println(findAnagrams("abab", "ab"))
	//fmt.Println(findAnagrams("baa", "aa"))
	fmt.Println(findAnagrams("abbbaa", "abba"))
}
