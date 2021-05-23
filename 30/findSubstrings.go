package main

func findSubstring(s string, words []string) []int {
	l := len(words)
	if l == 0 {
		return nil
	}
	step := len(words[0]) //步长
	steps := l * step     //总步数
	sLen := len(s)

	wordsMap := make(map[string]int)

	for i := 0; i < l; i++ {
		if len(words[i]) != step {
			return []int{}
		}
		wordsMap[words[i]]++
	}

	res := make([]int, 0)
	for i := 0; i < sLen-steps+1; i++ {
		subMaps := make(map[string]int)
		j := 0
		equal := true
		for j = 0; j < l; j++ {
			begin := i + step*j
			end := begin + step
			if _, ok := wordsMap[s[begin:end]]; !ok {
				break
			} else {
				subMaps[s[begin:end]]++
				if subMaps[s[begin:end]] > wordsMap[s[begin:end]] {
					equal = false
					break
				}
			}
		}

		if !equal || j != l {
			continue
		} else {
			res = append(res, i)
		}
	}
	return res

}

func main() {
	//fmt.Println(findSubstring("bbbbbbbbbarfoothefoobarman", []string{"bar", "foo"}))
	//fmt.Println(findSubstring("wordgoodgoodgoodbestword", []string{"word", "good", "best"}))
	//fmt.Println(findSubstring("", []string{}))
	//fmt.Println(findSubstring2("bbbbbbbbbarfoothefoobarman", []string{"bar", "foo"}))
	//fmt.Println(findSubstring2("wordgoodgoodgoodbestword", []string{"word", "good", "best"}))
	//fmt.Println(findSubstring2("", []string{}))

}
