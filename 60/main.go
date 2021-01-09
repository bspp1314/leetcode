package main

import "fmt"

var int2Byte = map[int]byte{
	1: '1',
	2: '2',
	3: '3',
	4: '4',
	5: '5',
	6: '6',
	7: '7',
	8: '8',
	9: '9',
}

func getPermutation(n int, k int) string {

	fac := make([]int, n+1)
	fac[0] = 1
	can := make([]int, n)
	for i := 1; i <= n; i++ {
		fac[i] = fac[i-1] * i
		can[i-1] = i
	}

	k = k - 1
	var res []byte
	for i := n - 1; i >= 0; i-- {
		index := k / fac[i]
		res = append(res, int2Byte[can[index]])
		can = append(can[0:index],can[index+1:]...)
		k -= index * fac[i]

	}

	return string(res)

}

func main() {
	fmt.Println(getPermutation(5, 2))
}
