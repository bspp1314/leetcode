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

	n1 := 0
	n2 := 0
	n3 := 0
	dp := make([]int,n)
	dp[0] = 1

	for i := 1; i < n; i++ {
		v1 := dp[n1] * 2
		v2 := dp[n2] * 3
		v3 := dp[n3] * 5

		dp[i] = Min(v1,Min(v2,v3))

		if v1 == dp[i] {
			n1++
		}

		if v2 == dp[i] {
			n2++
		}

		if v3 == dp[i] {
			n3++
		}

	}

	return dp[n-1]

}


func main() {
	fmt.Println(nthUglyNumber(10))
}
