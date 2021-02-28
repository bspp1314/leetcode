package main

import "fmt"

func titleToNumber(s string) int {
	res := 0

	for i := 0;i<len(s);i++ {
		res = res * 26 + (int(s[i] - 'A') + 1)
	}

	return res
}

func main()  {
	fmt.Printf("%d",int('A'))
}
