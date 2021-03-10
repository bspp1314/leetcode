package main


func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	temp :=  make(map[byte]int)

	for i := 0; i < len(s); i++ {
		temp[s[i]]++
		temp[t[i]]--
	}

	for _, v := range temp {
		if v != 0 {
			return false
		}

	}

	return true
}
