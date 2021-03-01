package main

import "fmt"

func countPrimes(n int) int {

	res := 0

	primes := make(map[int]bool)
	for i := 2; i < n; i++ {
		_,exist := primes[i]
		if exist {
			continue
		}

		res++
		for j := 0; i * j < n ; j++ {
			primes[i*j] = true
		}
	}
	return res
}

func main() {
	fmt.Println(countPrimes(10))
}
