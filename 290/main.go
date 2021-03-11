package main

import (
	"fmt"
	"strings"
)

func wordPattern(pattern string, s string) bool {
	words := strings.Split(s," ")
	if len(words) != len(pattern) {
		return false
	}

	kMap := make(map[byte]string)
	fmt.Println(words)

	for i:=0;i<len(pattern);i++ {
		word,exist := kMap[pattern[i]]
		if !exist {
			kMap[pattern[i]] = words[i]
		}else{
			if word != words[i] {
				return false
			}
		}
	}

	rewords := make([]string,len(pattern))
	for i := 0; i < len(pattern); i++ {
		rewords[i] = kMap[pattern[i]]
	}

	s = strings.Join(words,"")
	pattern = strings.Join(rewords,"")
	return s == pattern
}

func main() {
	fmt.Println(wordPattern("abba","dog cat cat dog"))
}