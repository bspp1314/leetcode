package main

import "fmt"

func numDistinct(s string, t string) int {
	tLen := len(t)
	sLen := len(s)
	dp := make([]int, tLen+1)
	dp[0] = 1

	tMap := make([]int, 128)
	for i := 0; i < 128; i++ {
		tMap[i] = -1
	}

	next := make([]int, tLen)

	for i := 0; i < tLen; i++ {
		index := int(t[i])
		next[i] = tMap[index]
		tMap[index] = i
	}

	for i := 0; i < sLen; i++ {
		index := int(s[i])
		for j := tMap[index]; j >= 0; j = next[j] {
			dp[j+1] += dp[j]
		}
	}
	return dp[len(t)]
}
func main() {
	//str := "dadklasfjdklsjfklasdjklfa"
	//for i := 0; i < len(str); i++ {
	//	m := int(str[i])
	//	fmt.Printf("%d-", m)
	//}
	fmt.Println(numDistinct("rabbbit", "rabbit"))
	fmt.Println("vim-go")
}
