package main

import (
	"fmt"
	"log"
)

func numDecodings(s string) int {
	if len(s) == 0 {
		return 0
	}

	if len(s) == 0 || ((!(s[0] >= '1' && s[0] <= '9') ) && s[0] != '*'){
		return 0
	}

	dp := make([]int, len(s)+1)
	dp[0] = 1
	if s[0] == '*' {
		dp[1] = 9
	}else{
		dp[1] = 1
	}

	for i := 2; i <= len(s); i++ {
		if s[i-1] == '0' && s[i-2] == '0' {
			return 0
		}

		if  (!(s[i-1] >= '0' && s[i-1] <= '9') ) && s[i-1] != '*' {
			return 0
		}
		// s[i]  和 s[0 ... i-1]
		if s[i-1] == '*' {
			dp[i] += (dp[i-1] * 9) % 100000007
		}else{
			if s[i-1] >= '1' && s[i-1] <= '9' {
				dp[i] += dp[i-1]  % 100000007
			}
		}

		log.Println("==========",dp[i])

		if s[i-1] == '*' {
			if s[i-2] == '*' {
				dp[i] += (dp[i-2] * 15) % 100000007
			}else{
				if s[i-2] == '1' || s[i-2] == '2' {
					dp[i] +=(dp[i-2] * 9) % 100000007
				}
			}
		}else{
			if s[i-2] == '*' {
				if s[i-1] == '0'  &&  s[i-1] <= '6' {
					dp[i] +=(dp[i-2] * 2 ) % 100000007
				}else if s[i-1] >= '7' && dp[i] <= '9' {
					dp[i] += dp[i-2] % 100000007
				}
			}
		}

	}
	return dp[len(s)]



}
func main() {
	fmt.Println(numDecodings("**"))
}
