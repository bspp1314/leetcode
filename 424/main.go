package main

import (
	"fmt"
)

func characterReplacement(s string, k int) int {
	if len(s) == 0 || len(s) == 1 {
		return len(s)
	}

	maxFun := func(a int, b int) int {
		if a > b {
			return a
		} else {
			return b
		}
	}

	slow := 0
	fast := 0

	windows := [256]int{}
	highestFreChar := 0

	for fast < len(s) {
		windows[s[fast]]++
		highestFreChar = maxFun(highestFreChar, windows[s[fast]])
		for highestFreChar+k < fast-slow+1 {
			windows[s[slow]]--
			slow++
		}
		fast++
	}


	return fast - slow

}

func main() {
	out := characterReplacement("AAAABBB", 0)
	fmt.Println("out", out)
}
