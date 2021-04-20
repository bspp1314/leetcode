package main

import (
	"fmt"
	"math"
)

func nthSuperUglyNumber(n int, primes []int) int {
	dp := make([]int,n)
	dp[0] = 1

	primesLen := len(primes)
	primesIndex := make([]int,primesLen)

	for i := 1; i < n; i++ {

		min := math.MaxInt32
		var minIndex []int

		for j := 0; j < primesLen; j++ {
			v := dp[primesIndex[j]]	 * primes[j]
			if v < min {
				min = v
				minIndex = []int{j}
			}else if v == min {
				minIndex = append(minIndex,j)
			}
		}


		dp[i] = min

		for k := 0; k < len(minIndex); k++ {
			primesIndex[minIndex[k]]++
		}

	}

	return dp[n-1]
}

func main() {
	fmt.Println(nthSuperUglyNumber(10,[]int{2,3,5}))
}
