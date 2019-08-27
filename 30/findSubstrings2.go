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
	for i := 0; i < stepSize; i++ {
		subMap := make(map[string]int)
		for j := i; j < sLen-stepSize; j += stepSize {

		}
	}

}

func main() {
	//fmt.Println(findSubstring("barfoothefoobarman", []string{"bar", "foo"}))
	//fmt.Println(findSubstring("wordgoodgoodgoodbestword", []string{"word", "good", "best"}))
	//fmt.Println(findSubstring("", []string{}))
	//fmt.Println(findSubstring2("bbbbbbbbbarfoothefoobarman", []string{"bar", "foo"}))
	//fmt.Println(findSubstring2("wordgoodgoodgoodbestword", []string{"word", "good", "best"}))
	//fmt.Println(findSubstring2("", []string{}))

}
