package main

import (
	"fmt"
	"strconv"
)

func ByteToInt(b byte) int {
	return int(b - '0')
}

func IntToByte(i int) byte {
	return byte(i + '0')
}


func addStrings(num1 string, num2 string) string {
	index := 0
	ans := ""
	for i, j := len(num1) - 1, len(num2) - 1; i >= 0 || j >= 0 || index != 0; i, j = i - 1, j - 1 {
		var x, y int
		if i >= 0 {
			x = int(num1[i] - '0')
		}
		if j >= 0 {
			y = int(num2[j] - '0')
		}
		result := x + y + index
		ans = strconv.Itoa(result%10) + ans
		index = result / 10
	}

	return ans
}

func multiply(num1 string, num2 string) string {
	if num1 == "0" || num2 == "0" {
		return "0"
	}
	m := len(num1)
	n := len(num2)
	results := make([]int, m+n)
	for i := m - 1; i >= 0; i-- {
		x := int(num1[i]) - '0'
		for j := n - 1; j >= 0; j-- {
			y := int(num2[j] - '0')
			results[i+j+1] += x * y
		}
	}

	for i := m + n - 1; i > 0; i-- {
		results[i-1] += results[i] / 10
		results[i] %= 10
	}


	result := ""
	idx := 0
	if results[0] == 0 {
		idx = 1
	}
	for ; idx < m+n; idx++ {
		result += strconv.Itoa(results[idx])
	}
	return result

}


func main() {
	fmt.Println(multiply("A", "A"))
	fmt.Println(ByteToInt('1'))
}
