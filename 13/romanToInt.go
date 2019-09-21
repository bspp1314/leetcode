package main

import (
	"fmt"
)

func romanToInt(s string) int {
	romanMap := map[byte]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}
	l := len(s)
	if l == 1 {
		return romanMap[s[0]]
	}
	res := romanMap[s[0]]

	for i := 1; i < l; i++ {
		if romanMap[s[i-1]] < romanMap[s[i]] {
			res = res - romanMap[s[i-1]] + (romanMap[s[i]] - romanMap[s[i-1]])
		} else {
			res += romanMap[s[i]]
		}
	}
	return res
}
func main() {
	fmt.Println(romanToInt("MMM"))
	fmt.Println(romanToInt("MMMIX"))
	fmt.Println(romanToInt("MMMXCIX"))
	fmt.Println(romanToInt("MMMCMXCIX"))
}
