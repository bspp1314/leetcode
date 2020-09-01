package main

import "fmt"

func checkInclusion(s1 string, s2 string) bool {
	cMap := [26]int{}

	for i := 0; i < len(s1); i++ {
		cMap[s1[i] - 'a']++
	}

	right := 0
	left := 0

	for right < len(s2) {
		//已存在
		if cMap[s2[right] - 'a'] <= 0 {
			for s2[left] != s2[right] {
				cMap[s2[left] - 'a']++
				left++
			}
			left++
		}else{
			cMap[s2[right] - 'a'] --
		}

		if (right - left + 1) == len(s1) {
			return true
		}
		right++
	}

	return false
}

func main() {
	fmt.Println(checkInclusion("", "abceabc"))
}
