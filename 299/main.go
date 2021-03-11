package main

import "fmt"

// 1 公牛和 3 奶牛。公牛是 8，奶牛是 0, 1 和 7。
func getHint(secret string, guess string) string {
	secretMap := make(map[byte]int)
	var subGuess []byte
	A := 0
	i := 0
	j := 0
	n := len(secret)
	m := len(guess)
	for i < n && j < m {
		if secret[i] == guess[j] {
			A++
		}else{
			secretMap[secret[i]]++
			subGuess = append(subGuess,guess[j])
		}

		i++
		j++
	}

	for i < n  {
		secretMap[secret[i]]++
		i++
	}

	fmt.Println(secretMap)

	B := 0
	for i := 0; i < len(subGuess); i++ {
		v,_ := secretMap[subGuess[i]]
		if v > 0 {
			B++
			secretMap[subGuess[i]]--
		}
	}

	for j < m  {
		v,_ := secretMap[guess[j]]
		if v > 0 {
			B++
			secretMap[guess[j]]--
		}
		j++
	}

	res := ""
	if A != 0 {
		res = fmt.Sprintf("%dA",A)
	}

	if B != 0 {
		res += fmt.Sprintf("%dB",B)
	}
	return res
}

func main() {
	fmt.Println(getHint("1123","0111"))
}
