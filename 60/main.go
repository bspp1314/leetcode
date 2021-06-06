package main

import "fmt"


func getPermutation(n int, k int) string {
	factorial := make([]int, n+1)
	factorial[0] = 1
	candidates := make([]int,n)
	for i := 1; i <= n; i++ {
		factorial[i] = factorial[i - 1] * i
		candidates[i-1] = i
	}
	k--
	fmt.Println(candidates)

	sb := make([]byte,0)
	for i := n-1;i >=0;i--{
		index := k / factorial[i]
		sb = append(sb,byte('0'+candidates[index]))
		candidates = append(candidates[0:index],candidates[index+1:]...)
		k = k - index * factorial[i]
	}


	return string(sb)

}

func main() {
	fmt.Println(getPermutation(5, 2))
}
