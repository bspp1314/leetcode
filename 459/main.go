package main

import "fmt"

func repeatedSubstringPattern(s string) bool {
	if len(s) <= 1 {
		return false
	}

	n := len(s)
	n2 := n >> 1

	for i := 1; i <= n2; i++ {
		subLen := i
		if n%subLen != 0 {
			continue
		}

		if sub(s[:i], subLen, s) {
			return true
		}
	}

	return false
}

func sub(sub string, nSub int, parent string) bool {

	n := len(parent)

	for i := 0; i < n; i += nSub {
		if sub != parent[i:i+nSub] {
			return false
		}
	}

	return true

}

func main() {
	fmt.Println(sub("ab", 2, "ababbc"))
	fmt.Println(repeatedSubstringPattern("abcabcabcabc"))
	fmt.Println(repeatedSubstringPattern("abab"))
	fmt.Println(repeatedSubstringPattern("aba"))
	fmt.Println(repeatedSubstringPattern("ab"))
}
