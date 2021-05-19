package main

import (
	"fmt"
	"math"
)

func isPalindrom(x int) bool {
	if x < 0 {
		return false
	} else {
		res := 0
		for res < x {
			res = 10*res + x%10
			x = x / 10
		}

		if res == x {
			return true
		} else {
			res = res / 10
			if res == x {
				return true
			} else {
				return false
			}
		}
	}
}

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}

	if x % 10 == 0 {
		return false
	}

	t := x
	val := 0
	for x != 0 {
		if math.MaxInt32/10 < val  {
			return false
		}
		val = val * 10 + x % 10
		x = x / 10
	}
	return t == val
}

func main() {
	//fmt.Println(isPalindrom(0))
	//fmt.Println(isPalindrom(1243421))
	//fmt.Println(isPalindrom(123321))
	fmt.Println(isPalindrome(121))
}
