package main

import "fmt"

func grayCode(n int) []int {
	if n == 0 {
		return []int{0}
	}

	res := []int{0}

	for i:=1;i<=n ;i++ {
		k := make([]int,2 * len(res))
		copy(k,res)
		for j := 0; j < len(res); j++ {
			k[len(res)+j] = res[len(res)-1-j] << 1 + 1
		}

		res = k
	}


	return res
}

func Print(a []int)   {
	for i := 0; i < len(a); i++ {
		fmt.Printf("%03b\n",i)
	}
}

func main() {
	Print(grayCode(3))
}
