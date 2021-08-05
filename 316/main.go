package main



func removeDuplicateLetters(s string) string {
	if len(s) <= 1 {
		return s
	}

	left := [26]int{}
	for _, ch := range s {
		left[ch-'a']++
	}

	var stack []byte
	inStack := [26]bool{}

	//思路就是 遇到一个新字符 如果比栈顶小 并且在新字符后面还有和栈顶一样的 就把栈顶的字符抛弃了
	for i  := range s {
		ch := s[i]

		if !inStack[ch-'a'] {
			for len(stack) > 0 && ch < stack[len(stack)-1] {
				last := stack[len(stack)-1] - 'a'
				if left[last] == 0 {
					break
				}
				stack = stack[:len(stack)-1]
				inStack[last] = false
			}

			stack = append(stack, ch)
			inStack[ch-'a'] = true
		}

		left[ch-'a']--
	}

	return string(stack)


}
