package main

import (
	"fmt"
)

func longestPalindrome2(s string) string {
	if len(s) <= 1 {
		return s
	}

	dp := make([][]bool,len(s))
	for i:=0;i<len(s);i++ {
		dp[i] = make([]bool,len(s))
		dp[i][i] = true
	}

	max := 0
	maxBegin := 0
	maxEnd := 1
	for right := 0;right < len(s);right++ {
		for left := 0;left <= right;left++ {
			if left == right {
				continue
			}
			if s[left] == s[right] && (dp[left+1][right-1] || right == left+1)  {
				dp[left][right] = true
				newLen := right-left + 1

				if newLen > max {
					max = newLen
					maxBegin = left
					maxEnd = right+1
				}
			}
		}
	}



	return s[maxBegin:maxEnd]

}
func expendAroundCenter(s string, begin int, end int) int {
	l := len(s)
	for begin >= 0 && end < l && s[begin] == s[end] {
		begin--
		end++
	}
	return end - begin - 1
}

func longestPalindrome(s string) string {
	begin := 0
	end := 0
	l := len(s)

	if l <= 1 {
		return s
	}
	for i := 0; i < l; i++ {
		len1 := expendAroundCenter(s, i, i)
		len2 := expendAroundCenter(s, i, i+1)

		new_len := len1
		if len1 < len2 {
			new_len = len2
		}

		if new_len > (end - begin) {
			begin = i - (new_len-1)/2
			end = i + new_len/2
		}
	}
	return s[begin : end+1]
}

func min(x int, y int) int {
	if x > y {
		return y
	}
	return x
}
func longestPalindrome3(s string) string {
	l := len(s)
	s2 := "$#"
	for i := 0; i < l; i++ {
		s2 += string(s[i])
		s2 += "#"
	}
	p := make([]int, len(s2))
	id := 0
	mx := 0
	resId := 0
	resMax := 0
	for i := 0; i < len(s2); i++ {
		if mx > i {
			p[i] = min(p[2*id-i], mx-i)
		} else {
			p[i] = 1
		}

		for (i-p[i] >= 0) && (i+p[i] < len(s2)) && s2[i+p[i]] == s2[i-p[i]] {
			p[i]++
		}
		if mx < i+p[i] {
			mx = i + p[i]
			id = i
		}
		if resMax < p[i] {
			resMax = p[i]
			resId = i
		}
	}
	return s[resId-resMax : resMax-1]

}
func main() {
	fmt.Println(longestPalindrome3("kkkm"))
}
