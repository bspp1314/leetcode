package main

import "fmt"

var nthMap = map[int]bool {
	1:true,
	2:true,
	3:true,
	4:true,
	5:true,
	6:true,
	8:true,
	9:true,
	10:true,
}

func Min(a,b int) int   {
	if a < b {
		return a
	}

	return b
}

func nthUglyNumber(n int) int{
	if n == 0 {
		return 1
	}
	dp := make([]int,n)
	dp[0] = 1

	i2 := 0
	i3 := 0
	i5  := 0

	count := 1

	for count < n {
		m2 := dp[i2] * 2
		m3 := dp[i3] * 3
		m5 := dp[i5] * 5

		//get min
		min := Min(m2,Min(m3,m5))
		if min == m2 {
			i2++
		}else if min == m3 {
			i3++
		}else{
			i5++
		}

		dp[count] = min
		count++
	}

	return dp[count-1]

}


func main() {
	fmt.Println(nthUglyNumber(1000))
}
