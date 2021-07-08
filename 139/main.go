package main

func wordBreak(s string, wordDict []string) bool {
	if len(s) == 0 {
		return true
	}

	if len(wordDict) == 0 {
		return false
	}

	wordDictMap := make(map[string]bool)

	for i := 0;i < len(wordDict);i++ {
		wordDictMap[wordDict[i]] = true
	}


	//dp[i]

	dp := make([]bool, len(s) + 1)
	dp[0] = true

	for i := 1; i <= len(s);i++ {
		for j := 0;j < i;j++ {
			if dp[j] && wordDictMap[s[j:i]] {
				dp[i] = true
				break
			}
		}
	}

	return dp[len(s)]

}

func main() {
	
}
