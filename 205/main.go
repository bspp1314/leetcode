package main

import "fmt"

//输入：s = "egg", t = "add"
func isIsomorphic(s string, t string) bool {
	ls := len(s)
	lt := len(t)
	if ls != lt {
		return false
	}

	if ls <= 1 {
		return true
	}

	preIndexOfs := [256]int{}
	preIndexOft := [256]int{}

	//"bbbaaaba"
	//"aaabbbba"
	for i := 0; i < ls; i++ {
		if preIndexOfs[s[i]] != preIndexOft[t[i]] {
			return false
		}
		preIndexOfs[s[i]] = i + 1
		preIndexOft[t[i]] = i + 1
	}

	return true
}

func main() {
	fmt.Println(isIsomorphic("bbbaaaba", "aaabbbba"))

}
