package main

import (
	"fmt"
)

//https://www.cnblogs.com/grandyang/p/4401196.html
func isMatch(s, p string) bool {
	i := 0
	j := 0
	jStart := -1
	iStart := -1


	// 关于 * 号的局限
	//s中有任何字符串，星号都能匹配，功能很前大
	//也有其处理不了的问题，那就是一旦p中有s中不存在的字符，那么一定无法匹配，因为星号只能增加字符，不能消除字符，
	//再有就是星号一旦确定了要匹配的字符串，对于星号位置后面的匹配情况也就鞭长莫及了。
	for i < len(s) {
		if j < len(p) && (s[i] == p[j] || p[j] == '?') {
			i++
			j++
		} else if j < len(p) && p[j] == '*' {
			iStart = i
			jStart = j
			j = jStart + 1
		} else if jStart != -1 {
			// 当发现 s[i] 和 p[j] 无法匹配时，但是好在之前 p[jStart] 之前出现了星号，把 s[i] 交给 p[jStart] 的星号去匹配。
			j = jStart + 1
			iStart += 1
			i = iStart
		} else {
			return false
		}
	}

	//匹配完了s中的所有字符，但是之后还要检查p串，此时没匹配完的p串里只能剩星号，不能有其他的字符，将连续的星号过滤掉，如果j不等于p的长度
	for j < len(p) {
		if p[j] != '*' {
			return false
		}
		j++
	}

	return true
}

func main() {
	a := []int{2}
	a = append(a,10)

	fmt.Println(isMatch("aab", "c*a*b*"))
}
