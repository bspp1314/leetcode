package main

import "fmt"

func reverseVowels(s string) string {
	if len(s) <= 1 {
		return s
	}

	left := 0
	res := make([]byte,len(s))
	right := len(s) -1
	VowelsMap := map[byte]bool{
		'a':true,
		'e':true ,
		'i':true,
		'o':true,
		'u':true,
		'A':true,
		'E':true,
		'I':true,
		'O':true,
		'U':true,
	}

	for left <= right {
		if !VowelsMap[s[left]] {
			res[left] = s[left]
			left++
			continue
		}

		if !VowelsMap[s[right]] {
			res[right] = s[right]
			right--
			continue
		}
		res[left] = s[right]
		res[right] = s[left]
		left++
		right--
	}

	return  string(res)
}

func main()  {
	fmt.Println(reverseVowels("leetcode"))
}