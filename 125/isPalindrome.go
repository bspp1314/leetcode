package main

import "fmt"

func isBigLetter(c byte)bool  {
	return c >= 'A' && c <= 'Z'
}

func isSmallLetter(c byte)bool  {
	return c >= 'a' && c <= 'z'
}

func isLetter(c byte)bool  {
	return isBigLetter(c) || isSmallLetter(c)
}

func isPalindrome(a string)bool  {
	left := 0
	right := len(a) -1

	for left <= right {
		m := a[left]
		n := a[right]


		if !isLetter(m) {
			left ++
			continue
		}

		if !isLetter(n) {
			right--
			continue
		}


		if isBigLetter(m) {
			m = m + 32
		}

		if isBigLetter(n) {
			n = n + 32
		}

		if m == n {
			left++
			right--
		}else{
			return  false
		}
	}

	return true
}

func main() {
	out := isPalindrome("A man, a plan, a canal: Panama")
	fmt.Println(out)
}
