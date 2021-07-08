package main

import "fmt"

func reverseWords(s string) string {
	var res []byte
	left := 0
	right := len(s) - 1

	for left <= right && s[left] == ' ' {
		left++
	}

	for left <= right && s[right] == ' ' {
		right--
	}

	if left > right {
		return ""
	}

	end := right + 1
	for left <= right {
		if s[right] == ' ' {
			res = append(res, []byte(s[right+1:end])...)
			res = append(res, ' ')
			for left <= right && s[right] == ' ' {
				right--
				end = right + 1
			}
		} else {
			right--
		}
	}

	res = append(res, []byte(s[right+1:end])...)

	return string(res)

}

func main() {
	fmt.Println(reverseWords("the ske is blue"))
}
