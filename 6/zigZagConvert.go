package main

import "fmt"

func convert2(s string, numRows int) string {
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

/*
	LEET CODE ISHI RING   //4
	L      D      R   // 0 + 6    12
	E    O E    I I  //  1 5 7 11 13
	E  C   I  H   N      2 4 8 10  14
	T      S      G      3   9    15
*/

func convert3(s string, numRows int) string {
	if numRows == 0 || len(s) <= 1 {
		return s
	}

	resBytes := make([]byte, 0, len(s))
	line := 0
	n := len(s)

	for line < numRows {
		index := line
		z := true
		for index < n {
			resBytes = append(resBytes, s[index])
			if line == 0 || line == numRows-1 {
				index += 2*numRows - 2
			} else {
				if z {
					index += 2*numRows - 2 - line
					z = false
				} else {
					index += line
					z = true
				}
			}
		}
		line++
	}

	return string(resBytes)

}

func main() {
	fmt.Println(convert("PAYPALISHIRING", 4))
	fmt.Println(convert2("PAYPALISHIRING", 4))
	fmt.Println(convert3("PAYPALISHIRING", 4))
}
