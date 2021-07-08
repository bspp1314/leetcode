package main

import (
	"fmt"
	"log"
	"strings"
)


func wordBreak(s string, wordDict []string) (sentences []string) {
	wordSet := make(map[string]bool)
	for _, w := range wordDict {
		wordSet[w] = true
	}

	n := len(s)
	dp := make([][][]string, n+1)
	var backtrack func(index int) [][]string
	// 对于字符串 ss，如果某个前缀是单词列表中的单词，则拆分出该单词，然后对 s 的剩余部分继续拆分。

	backtrack = func(index int) [][]string {
		if dp[index] != nil {
			return dp[index]
		}

		var wordsList [][]string

		for i := index+1;i <= n;i++ {
			if wordSet[s[index:i]] {
				if i == n {
					wordsList = append(wordsList, []string{s[index:i]})
				}else{
					nextWordLists := backtrack(i)
					for _, list := range nextWordLists {
						wordsList = append(wordsList,append([]string{s[index:i]},list...))
					}
				}
			}
		}

		dp[index] = wordsList
		return wordsList
	}



	for _, words := range backtrack(0) {
		sentences = append(sentences, strings.Join(words, " "))
	}
	return
}


func main() {
	s := wordBreak("cat",[]string{"a","cat"})
	fmt.Println("s is ",s)
	for i:=0;i<len(s);i++ {
		log.Println(s[i])
	}

}