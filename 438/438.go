package main

import "fmt"

func findAnagrams(s string, p string) []int {
	sLen := len(s)
	pLen := len(p)
	if sLen == 0 || pLen == 0 {
		return nil
	}
	pMap := make([]int, 256)
	window := make([]int, 256)
	cn := 0
	for i := 0; i < pLen; i++ {
		if pMap[p[i]] == 0 {
			cn++
		}
		pMap[p[i]]++
	}
	left := 0
	right := 0
	formed := 0
	res := make([]int, 0)

	for right < sLen {
		if pMap[s[right]] == 0 {
			window = make([]int, 256)
			formed = 0
			right++
			left = right
		} else {
			if window[s[right]] == 0 {
				formed++
			}
			if right == 2 {
				fmt.Println("left", left)
				fmt.Println("right", left)
				fmt.Println("formed", formed)
				fmt.Println("cn", cn)
			}
			window[s[right]]++
			if window[s[right]] <= pMap[s[right]] {
				if formed == cn {
					window[s[left]]--
					formed--
					res = append(res, left)
					right++
					left++
				} else { //formed < cn
					right++
				}
			} else { //window[s[right]] > pMap[s[right]]
				window[s[left]]--
				left++
				right++

			}
		}
	}
	return res

}

func main() {
	//fmt.Println(findAnagrams("cbac", "bca"))
	fmt.Println(findAnagrams("aabcaa", "abac"))
	//fmt.Println(findAnagrams("cbaebabacd", "bca"))
	//fmt.Println(findAnagrams("abab", "ab"))
}
