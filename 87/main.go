package main

var Memory  = make(map[[2]string]bool)
func isScramble(s1 string, s2 string) bool {
	if Memory[[2]string{s1,s2}] {
		return true
	}

	if s1 == s2 {
		return true
	}

	if len(s1) != len(s2) {
		return false
	}

	if len(s1) == 1 {
		return s1 == s2
	}

	byteMap := make(map[byte]int, 0)

	for i := 0; i < len(s1); i++ {
		byteMap[s1[i]] = byteMap[s1[i]] + 1
		byteMap[s2[i]] = byteMap[s2[i]] - 1
	}

	for _, i := range byteMap {
		if i != 0 {
			return false
		}
	}

	n := len(s1)
	for i := 1; i < n; i++ {
		if isScramble(s1[:i], s2[:i]) && isScramble(s1[i:len(s1)], s2[i:len(s1)]) {
			Memory[[2]string{s1,s2}] = true
			return true
		}

		if isScramble(s1[:i], s2[n-i:]) && isScramble(s1[i:], s2[0:n-i]) {
			Memory[[2]string{s1,s2}] = true
			return true
		}
	}

	return false

}

func main() {
	Memory[[2]string{"great","eatgr"}] = true
}
