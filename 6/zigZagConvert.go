package main

import "fmt"

func convert(s string, numRows int) string {
	if len(s) <= 1 || numRows == 1 {
		return s
	}

	res := make([]byte, len(s))
	inter1 := 2*numRows - 2
	m := 0
	for i := 0; i < numRows; i++ {
		if i == 0 || i == numRows-1 {
			for j := i; j < len(s); j += inter1 {
				res[m] = s[j]
				m++
			}
		} else {
			inter2 := 2 * i
			for j := i; j < len(s); j += inter2 {
				res[m] = s[j]
				m++
				inter2 = inter1 - inter2
			}
		}
	}
	return string(res)
}
func convert(s string, numRows int) string {
	if len(s) <= 1 || numRows == 1 {
		return s
	}

	res := make([]byte, len(s))
	inter1 := 2*numRows - 2
	m := 0
	for i := 0; i < numRows; i++ {
		inter2 := 2 * i
		for j := i; j < len(s); j += inter2 {
			res[m] = s[j]
			m++
			if i == 0 || i == numRows-1 {
				inter2 = inter1
			} else {
				inter2 = inter1 - inter2
			}
		}
	}
	return string(res)
}
func main() {
	fmt.Println(convert("PAYPALISHIRING", 4))
	fmt.Println(convert2("PAYPALISHIRING", 4))
}
