package main

import (
	"fmt"
)

//暴力法
var countMap map[string]bool = map[string]bool{
	"1":  true,
	"2":  true,
	"3":  true,
	"4":  true,
	"5":  true,
	"6":  true,
	"7":  true,
	"8":  true,
	"9":  true,
	"10": true,
	"11": true,
	"12": true,
	"13": true,
	"14": true,
	"15": true,
	"16": true,
	"17": true,
	"18": true,
	"19": true,
	"20": true,
	"21": true,
	"22": true,
	"23": true,
	"24": true,
	"25": true,
	"26": true,
}

func numDecodings(s string) int {
	if len(s) <= 1 {
		if countMap[s] {
			return 1
		}
		return 0
	}

	var n1 int
	var n2 int
	if !countMap[s[:1]] {
		n1 = 0
	} else {
		n1 = numDecodings(s[1:])
	}
	if !countMap[s[:2]] {
		n2 = 0
	} else {
		n2 = numDecodings(s[2:])
	}
	return n1 + n2
}

func numDecodings2(s string) int {
	if len(s) == 0 || (!(s[0] >= '1' && s[0] <= '9')) {
		return 0
	}

	dp := make([]int, len(s)+1)
	dp[0] = 1
	dp[1] = 1
	for i := 2; i <= len(s); i++ {
		low := s[i-1] - '0'
		high := s[i-2] - '0'
		if low > 0 {
			dp[i] += dp[i-1]
		}
		sum := high*10 + low
		if sum >= 10 && sum <= 26 {
			dp[i] += dp[i-2]
		} else {
			if low == 0 {
				return 0
			}
		}
	}
	return dp[len(s)]
}
func MemNumDecodings(s string) int {
	dp := make([]int, len(s)+1)
	for i := 0; i <= len(s); i++ {
		dp[i] = -1
	}
	return MemNumDecodingsHelp(s, dp)
}
func MemNumDecodingsHelp(s string, dp []int) int {
	if dp[len(s)] >= 0 {
		return dp[len(s)]
	}
	if len(s) <= 1 {
		if countMap[s] {
			dp[len(s)] = 1
		} else {
			dp[len(s)] = 0
		}
		return dp[len(s)]
	}
	var n1 int
	var n2 int
	if !countMap[s[:1]] {
		n1 = 0
	} else {
		n1 = MemNumDecodingsHelp(s[1:], dp)
	}
	if !countMap[s[:2]] {
		n2 = 0
	} else {
		n2 = MemNumDecodingsHelp(s[2:], dp)
	}
	dp[len(s)] = n1 + n2
	return dp[len(s)]
}



func numDecodings3(s string) int {
	if len(s) == 0 {
		return 0
	}

	if !(s[0] >= '1' && s[0] <= '9') {
		return 0
	}


	dp := make([]int,len(s)+1)

	dp[0] = 0
	dp[1] = 1



	for i := 2;i<=len(s);i++ {
		low := s[i-1] - '0'
		high := s[i-2] - '0'

		if low > 0 {
			dp[i] += dp[i-1]
		}
		v := high * 10 + low

		if v >= 10 && v <= 26 {
			dp[i] += dp[i-2]
		}else if low == 0 { // low is 0 and high is 0
			return 0
		}
	}

	return dp[len(s)]
}

func minimumTotal(triangle [][]int) int {
	if len(triangle) == 0 {
		return 0
	}

	if len(triangle) == 1 {
		return triangle[0][0]
	}

	res := triangle[0][0]
	for i := 1;i < len(triangle);i++ {
		var min int
		min = 1<<63 - 1


		for j:=0;j < len(triangle[i]);j++ {
			if j == 0 {
				triangle[i][j] += triangle[i-1][j]
			}else if j == len(triangle[i])-1 {
				triangle[i][j] += triangle[i-1][j-1]
			}else{
				v1 := triangle[i-1][j-1]
				v2 := triangle[i-1][j]
				if v1 < v2 {
					triangle[i][j] += v1
				}else{
					triangle[i][j] += v2
				}
			}

			if triangle[i][j] < min {
				min = triangle[i][j]
			}

		}
		if i == len(triangle) - 1 {
			res = min

		}
	}

	return res

}

func main() {
	fmt.Println(numDecodings3("20"))
	fmt.Println(numDecodings2("20"))
	//fmt.Println(numDecodings("12"))
	//fmt.Println(numDecodings2("01"))
	//fmt.Println(numDecodings2("1212368242421412121"))
	//fmt.Println(MemNumDecodings("1212368242421412121"))
	//fmt.Println(MemNumDecodings("1299999999999"))
}
