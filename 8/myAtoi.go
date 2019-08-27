package main

import (
	"fmt"
	"math"
	"strconv"
)

func myAtoi(str string) int {
	sLen := len(str)
	if sLen == 0 {
		return 0
	}
	symbol := false
	res := int64(0)
	active := true
	for i := 0; i < sLen; i++ {
		ch := str[i] - '0'
		if ch >= 0 && ch <= 9 {
			if active {
				res = res*10 + int64(ch)
			} else {
				res = res*10 - int64(ch)
			}
			symbol = true
			if res >= math.MaxInt32 {
				return math.MaxInt32
			}
			if res <= math.MinInt32 {
				return math.MinInt32
			}
		} else {
			if symbol {
				return int(res)
			} else {
				if str[i] == '-' {
					active = false
					symbol = true
				} else if str[i] == '+' {
					symbol = true
				} else if str[i] == ' ' {
					continue
				} else {
					return int(res)
				}
			}
		}
	}
	return int(res)
}

func main() {
	fmt.Println("vim-go")
	fmt.Println(strconv.Atoi("   333"))
	fmt.Println(myAtoi("   -+333"))
}
