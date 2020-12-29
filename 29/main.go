package main

import (
	"math"
)

func divide(dividend int, divisor int) int {
	if dividend == 0 {
		return 0
	}

	if divisor == 1 {
		return dividend
	}


	if divisor == -1 {
		if dividend != math.MinInt32 {
			return -dividend
		}

		return math.MaxInt32
	}

	flag := (dividend ^ divisor) > 0

	if dividend < 0 {
		dividend = -dividend
	}

	if divisor < 0 {
		divisor = -divisor
	}


	if flag {
		return Div(dividend,divisor)
	}

	return -Div(dividend,divisor)
}

func Div(a,b int) int  {
	if a < b {
		return 0
	}

	count := 1
	newB := b

	for (newB << 1)  <= a   {
		count = count << 1
		newB = newB << 1
	}

	return count + Div(a-newB,b)
}




func main() {
}
