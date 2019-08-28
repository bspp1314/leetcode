package main

import (
	"fmt"
	"sort"
	"strings"
)

type Bytes []byte

func (b Bytes) Swap(i, j int) {
	b[i], b[j] = b[j], b[i]
}
func (b Bytes) Less(i, j int) bool {
	return b[i] < b[j]
}
func (b Bytes) Len() int {
	return len(b)
}

func groupAnagrams(s []string) [][]string {
	l := len(s)
	if l == 0 {
		return nil
	}
	groupsMap := make(map[string][]string)
	for i := 0; i < l; i++ {
		m := Bytes(s[i])
		sort.Sort(m)
		k := string(m)

		strs := groupsMap[k]
		if len(strs) == 0 {
			strs = make([]string, 0)
		}
		strs = append(strs, s[i])
		groupsMap[k] = strs
	}
	res := make([][]string, 0, l)
	for _, v := range groupsMap {
		res = append(res, v)
	}
	return res
}

func groupAnagrams2(s []string) [][]string {
	l := len(s)
	if l == 0 {
		return nil
	}
	groupsMap := make(map[string][]string)
	for i := 0; i < l; i++ {
		byteArray := make([]byte, 26)
		l2 := len(s[i])
		for j := 0; j < l2; j++ {
			byteArray[s[i][j]-'a']++
		}
		var b strings.Builder
		for j := 0; j < 26; j++ {
			b.WriteByte(byteArray[j])
		}
		k := b.String()

		strs := groupsMap[k]
		if len(strs) == 0 {
			strs = make([]string, 0)
		}
		strs = append(strs, s[i])
		groupsMap[k] = strs
	}
	res := make([][]string, 0, l)
	for _, v := range groupsMap {
		res = append(res, v)
	}
	return res
}
func main() {
	fmt.Println(groupAnagrams([]string{"ate", "eta", "tae", "nat", "tan", "", ""}))
	fmt.Println(groupAnagrams2([]string{"ate", "eta", "tae", "nat", "tan", "", ""}))
	fmt.Println("vim-go")
	var b strings.Builder
	b.WriteByte('0')
	b.WriteByte('1')
	b.WriteByte('2')
	fmt.Println(b.String())
}
