package main

import (
	"fmt"
	"math"
)

func Abs(a int) int {
	if a > 0 {
		return a
	}

	return -a
}



func mySqrt(x float64) float64 {
	if x == 0 || x == 1 {
		return 1
	}

	result := x
	lastValue := 0.0

	for {
		lastValue = result
		result = result / 2.0 + x / 2.0 / result
		if math.Abs(result-lastValue)  < 0.99  {
			break
		}
	}

	return result
}

func main() {
	fmt.Println(mySqrt(9))
	fmt.Println(mySqrt(15))

}
