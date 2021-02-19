package main

import (
	"fmt"
	"math"
)

func mySqrt(x int) int {
	left, right := 0, x
	ans := -1
	for left <= right {
		mid := left + (right-left)/2
		if mid*mid <= x {
			ans = mid
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return ans
}
func main() {
	for i := 0; i <=10000 ; i++ {
		if !(int(math.Sqrt(float64(i))) == mySqrt(i)) {
			fmt.Println("i",i)
		}

	}
}
