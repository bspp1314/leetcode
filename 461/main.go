package main

import "fmt"

// 1 1 0  1
// 0 1 0  1
// 0 1 0  1
func Max(a, b int) int {
	if a > b {
		return a
	}

	return b
}

func hammingDistance(x int, y int) int {
	v := x ^ y
	res := 0
	k := 0

	for k <= 31 {
		res += v & 1
		v = v >> 1
		k++
	}

	return res
}

func hammingDistance2(x int, y int) int {
	v := x ^ y
	res := 0
	k := 0

	for k <= 31 {
		res += v & 1
		v = v & (v - 1)
		k++
	}

	return res
}

func main() {
	fmt.Println(hammingDistance(1, 4))
}
