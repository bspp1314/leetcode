package main

import (
	"fmt"
	"runtime"
)

// 1 公牛和 3 奶牛。公牛是 8，奶牛是 0, 1 和 7。
func getHint(secret string, guess string) string {
	runtime.GC()
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

//求交集
func intersect(slice1, slice2 []string) []string {
	m := make(map[string]int)
	nn := make([]string, 0)
	for _, v := range slice1 {
		m[v]++
	}

	for _, v := range slice2 {
		times, _ := m[v]
		if times == 1 {
			nn = append(nn, v)
		}
	}
	return nn
}

func main() {
	out := intersect([]string{"1","1","1"},[]string{"1"})
	fmt.Println(out)
	fmt.Println(getHint("1123","0111"))
}
