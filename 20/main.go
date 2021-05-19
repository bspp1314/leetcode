package main

import "fmt"

var m1 = map[byte]byte {
	'{':'}',
	'(':')',
	'[':']',
}

var m2 = map[byte]byte {
	'}':'{',
	']':'[',
	')':'(',
}


func isValid(s string) bool {
	stack := make([]byte,0)


	for i:=0;i<len(s);i++ {
		if v,ok := m1[s[i]];ok {
			stack = append(stack,v)
		}else if _,ok := m2[s[i]];ok {
			if len(stack) == 0 {
				return false
			}

			if stack[len(stack)-1] != s[i] {
				return false
			}
			stack = stack[:len(stack)-1]

		}else{
			return false
		}
	}

	if len(stack) != 0 {
		return false
	}

	return true
}

func main() {
	fmt.Println(isValid("()"))
	//fmt.Println(isValid("{{{(((())))}}"))
}