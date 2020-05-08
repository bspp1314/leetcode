package main

import "fmt"

func plusOne(digits []int) []int {
	if len(digits) == 0 {
		return digits
	}

	if len(digits) == 1 {
		if digits[0] == 9 {
			return []int{1, 0}
		} else {
			digits[0] += 1
			return digits
		}
	}

	if digits[len(digits)-1] != 9 {
		digits[len(digits)-1] += 1
		return digits
	} else {
		res := plusOne(digits[:len(digits)-1])
		res = append(res, 0)
		return res
	}

}

func plusOne2(digits []int) []int {
	if len(digits) == 0 {
		return digits
	}

	for i := len(digits) - 1; i >= 0; i-- {
		if digits[i] != 9 {
			digits[i] += 1
			return digits
		}
		digits[i] = 0

	}

	return append([]int{1}, digits...)
}

func main() {
	fmt.Println(plusOne2([]int{1, 2, 3, 4}))
	fmt.Println(plusOne2([]int{1, 2, 3, 9}))
	fmt.Println(plusOne2([]int{9, 9, 9, 9}))
	fmt.Println("vim-go")
}
