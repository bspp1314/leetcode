package main

import "fmt"

func Calcu(a byte, b byte, c byte) int {
	newA := int(a - '0')
	newB := int(b - '0')
	if c == '*' {
		return newA * newB
	} else if c == '+' {
		return newA + newB
	} else {
		return newA - newB
	}
}

func diffWaysToCompute(input string) []int {
	l := len(input)
	if l&1 == 0 {
		return nil
	}
	return diffWaysToComputeHelp(0, l-1, input)
}
func diffWaysToComputeHelp(start, end int, input string) []int {
	if start == end {
		v := int(input[start] - '0')
		return []int{v}
	}
	if (end - start) == 2 {
		v := Calcu(input[start], input[end], input[start+1])
		return []int{v}
	}
	res := make([]int, 0)
	for i := start + 1; i < end; i += 2 {
		left := diffWaysToComputeHelp(start, i-1, input)
		right := diffWaysToComputeHelp(i+1, end, input)
		fmt.Println("i", i)
		fmt.Println("left", left)
		fmt.Println("right", right)
		for j := 0; j < len(left); j++ {
			for k := 0; k < len(right); k++ {
				switch input[i] {
				case '-':
					res = append(res, left[j]-right[k])
					break
				case '+':
					res = append(res, left[j]+right[k])
					break
				case '*':
					res = append(res, left[j]*right[k])
					break
				default:
					panic("error ")
				}
			}
		}
	}
	return res
}
func main() {
	fmt.Println(diffWaysToCompute("2*3-4"))
}
