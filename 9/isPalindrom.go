package main

import "fmt"

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

func main() {
	fmt.Println(isPalindrom(0))
	fmt.Println(isPalindrom(1243421))
	fmt.Println(isPalindrom(123321))
	fmt.Println(isPalindrom(-123321))
}
