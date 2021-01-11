package main

import "fmt"

func minCut(s string) int {
	if len(s) <= 1 {
		return 0
	}

	dp := make([][]bool,len(s))
	for i:=0;i<len(s);i++ {
		dp[i] = make([]bool,len(s))
		dp[i][i] = true
	}

	for right := 0; right < len(s); right++ {
		for left := 0; left <=right ; left++ {
			if s[left] == s[right]  && (right -left <= 2 || dp[left+1][right-1]){
				dp[left][right] = true
			}
		}
	}

	dp2 := make([]int,len(s)+1)



	for right := 0; right < len(s); right++ {
		dp2[right+1] = len(s)
		for left := 0; left <= right; left++ {
			if dp[left][right] {
				dp2[right + 1] = Min(dp2[left] + 1, dp2[right + 1])
			}
		}
	}


	return dp2[len(s)]
}

func Min(a,b int) int   {
	if a < b {
		return a
	}

	return b
}



func main() {
	fmt.Println(minCut("apjesgpsxoeiokmqmfgvjslcjukbqxpsobyhjpbgdfruqdkeiszrlmtwgfxyfostpqczidfljwfbbrflkgdvtytbgqalguewnhvvmcgxboycffopmtmhtfizxkmeftcucxpobxmelmjtuzigsxnncxpaibgpuijwhankxbplpyejxmrrjgeoevqozwdtgospohznkoyzocjlracchjqnggbfeebmuvbicbvmpuleywrpzwsihivnrwtxcukwplgtobhgxukwrdlszfaiqxwjvrgxnsveedxseeyeykarqnjrtlaliyudpacctzizcftjlunlgnfwcqqxcqikocqffsjyurzwysfjmswvhbrmshjuzsgpwyubtfbnwajuvrfhlccvfwhxfqthkcwhatktymgxostjlztwdxritygbrbibdgkezvzajizxasjnrcjwzdfvdnwwqeyumkamhzoqhnqjfzwzbixclcxqrtniznemxeahfozp"))
}

