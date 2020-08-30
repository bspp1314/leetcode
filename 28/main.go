package main

import "fmt"

func strStr(haystack string, needle string) int {
	if len(needle) == 0 {
		return 0
	}

	if len(needle) > len(haystack) {
		return -1
	}

	j := 0
	i := 0

	// h e l l o
	for i < len(haystack)-(len(needle)-1) {
		if haystack[i] != needle[j] {
			i++
			continue
		}

		v := i
		fmt.Println("bef i",i)
		fmt.Println("j",j)
		for j = 0; j < len(needle); j++ {
			if haystack[i] == needle[j] {
				i++
				continue
			}
		}
		fmt.Println("after ==============",i)
		fmt.Println("==============",j)
		fmt.Println("==============",i)


		if j == len(needle) {
			return v
		} else {
			j = 0
		}
	}

	return -1
}
func main() {
	fmt.Println(strStr("mississippi", "issip"))
	return
}
