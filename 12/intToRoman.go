package main

import "fmt"

func intToRoman(num int) string {
	d := [4][]string{
		[]string{"", "I", "II", "III", "IV", "V", "VI", "VII", "VIII", "IX"},
		[]string{"", "X", "XX", "XXX", "XL", "L", "LX", "LXX", "LXXX", "XC"},
		[]string{"", "C", "CC", "CCC", "CD", "D", "DC", "DCC", "DCCC", "CM"},
		[]string{"", "M", "MM", "MMM"},
	}
	return d[3][num/1000] +
		d[2][num/100%10] +
		d[1][num/10%10] +
		d[0][num%10]
}
func romanToInt(s string) int {
	romanMap := map[byte]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}
	l := len(s)
	if l == 1 {
		return romanMap[s[0]]
	}
	res := romanMap[s[0]]

	for i := 1; i < l; i++ {
		if romanMap[s[i-1]] < romanMap[s[i]] {
			res = res - romanMap[s[i-1]] + (romanMap[s[i]] - romanMap[s[i-1]])
		} else {
			res += romanMap[s[i]]
		}
	}
	return res
}

func main() {
	for i := 1; i <= 3999; i++ {
		j := romanToInt(intToRoman(i))
		if i != j {
			fmt.Println("===============")
		}
	}
	fmt.Println(intToRoman(3999))
	fmt.Println(intToRoman(3909))
	fmt.Println(intToRoman(3009))
	fmt.Println(intToRoman(3000))
}
