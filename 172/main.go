package main

// 5
// 5 10  15  20  25
// 1  1   1  1   2

func trailingZeroes(n int) int {
	if n < 5 {
		return 0
	}

	res := 0
	for n >= 5 {
		res += n / 5
		n = n / 5
	}

	return res

}

func main()  {

}
