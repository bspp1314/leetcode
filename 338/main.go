package main

import (
	"fmt"
)

func countBits(num int) []int {
	res := make([]int,num+1)
	if num ==  0 {
		return []int{0}
	}

	res[0] =0
	res[1] = 1
	k := 1
	for i:=2;i<= num;i++ {
		if i % k == 0 {
			k = k *  2
		}
		res[i] = res[i-k] + 1

	}

	return res

}

func main() {
	out := countBits(20)

	fmt.Println(out)
}
