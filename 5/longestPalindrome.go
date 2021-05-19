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

func longestPalindrome(s string) string {
	if len(s) <= 1 {
		return s
	}

	dp := make([][]bool,len(s))
	for i:=0;i<len(s);i++ {
		dp[i] = make([]bool,len(s))
		dp[i][i] = true
	}
	begin := 0
	end := 0


	for j:=0;j < len(s);j++ {
		for i:=0;i <=j;i++ {
			if (j - i) <= 1 || (s[i] == s[j] && dp[i+1][j-1]) {
				dp[i][j] = true
				if j -i > end -begin {
					end = j
					begin = i
				}
			}
		}
	}

	return s[begin:end+1]




}




func expand(s string,left int,right int,n int ) (int,int)  {
	for left >= 0 && right <= n {
		if s[left] != s[right] {
			break
		}
		left--
		right++
	}

	return left+1,right-1
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
	fmt.Println(longestPalindrome("kkkm"))
}
