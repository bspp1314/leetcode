package main

import (
	"fmt"
	"math"
)

func isSql(N int) (int, bool) {
	n := int(math.Sqrt(float64(N)))
	for i := 2; i < n; i++ {
		m := N % i
		if m == 0 {
			return m, false
		}
	}
	return 0, true
}
func diviorGame(N int) bool {
	if N <= 1 {
		return false
	}
	var m int
	var ok bool
	m, ok = isSql(N)
	if ok {
		return true
	}

	_, ok = isSql(m)
	if ok {
		return true
	}
	return false

}
func main() {
	fmt.Println("vim-go")
}
