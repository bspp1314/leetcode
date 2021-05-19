package main

import "strings"

func max (a,b int)  int  {
	if a > b {
		return a
	}

	return b
}

//暴力法
func longestSubstring2(s string,k int) int   {
	start := 0
	res := 0

	for start = 0; start < len(s);start++ {
		mask := 0

		cnt := [26]int{}

		for end := start; end < len(s); end++ {
			c := s[end] - 'a'
			cnt[c]++
			if cnt[c] >= k{
				mask &= ^(1 << c )
			}else{
				mask |= 1 <<c
			}

			if mask != 0 {
				res = max(res,end-start+1)
			}
		}
	}

	return res
}

func longestSubstring(s string, k int) int {
	if k <= 1 {
		return len(s)
	}

	cnt := [26]int{}
	for _, ch := range s {
		cnt[ch-'a']++
	}

	var split byte
	for i, c := range cnt {
		if c > 0 && c < k {
			split = 'a' + byte(i)
			break
		}
	}

	if split == 0 {
		return len(s)
	}


	ans := 0
	subStrs := strings.Split(s,string(split))
	for _, subStr := range subStrs {
		ans = max(ans, longestSubstring(subStr, k))
	}
	return ans

}

func main() {

}
