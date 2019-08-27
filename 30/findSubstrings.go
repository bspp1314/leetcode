package main

func findSubstring(s string, words []string) []int {
	l := len(words)
	if l == 0 {
		return nil
	}
	stepSize := len(words[0])
	steps := l * stepSize
	sLen := len(s)

	wordsMap := make(map[string]int)
	for i := 0; i < l; i++ {
		wordsMap[words[i]]++
	}

	res := make([]int, 0)
	for i := 0; i < sLen-steps+1; i++ {
		subMaps := make(map[string]int)
		j := 0
		for j = 0; j < l; j++ {
			begin := i + stepSize*j
			end := begin + stepSize
			keyStr := s[begin:end]
			if _, ok := wordsMap[keyStr]; !ok {
				goto Loop
			} else {
				subMaps[keyStr]++
				if subMaps[keyStr] > wordsMap[keyStr] {
					goto Loop
				}
			}
		}
		if j != steps {
			continue
		}

		for k, v := range wordsMap {
			if v2 := wordsMap[k]; v2 != v {
				goto Loop
			}
		}
		res = append(res, i)
	Loop:
		continue
	}
}

func main() {
	//fmt.Println(findSubstring("bbbbbbbbbarfoothefoobarman", []string{"bar", "foo"}))
	//fmt.Println(findSubstring("wordgoodgoodgoodbestword", []string{"word", "good", "best"}))
	//fmt.Println(findSubstring("", []string{}))
	//fmt.Println(findSubstring2("bbbbbbbbbarfoothefoobarman", []string{"bar", "foo"}))
	//fmt.Println(findSubstring2("wordgoodgoodgoodbestword", []string{"word", "good", "best"}))
	//fmt.Println(findSubstring2("", []string{}))

}
