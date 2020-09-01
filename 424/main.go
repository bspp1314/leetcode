package main

import "fmt"

func characterReplacement(s string, k int) int {
	if len(s) == 0 || len(s) == 1{
		return len(s)
	}

	j := 0
	cMap := make(map[byte]int)
	for i := 0;i < len(s);i++ {
		cMap[s[i]] ++
	}
	return res

}

func main() {
	out := characterReplacement("AAAA",0)
	fmt.Println("out",out)
	out2 := characterReplacement("ABACDE",1)
	fmt.Println("out",out2)
}
