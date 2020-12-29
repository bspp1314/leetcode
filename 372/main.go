package main

import (
	"fmt"
)


// (a * b) % k = (a % k)(b % k) % k
// (a * b) % k = ((a1k + a2) * (b1k + b2)) % k =  (a1b1k^2 +a1b2k+b1a2k+b1b2 ) % k = (a2b2)%k
// (a % k)(b % k) % k= ((a1k + a2)%k + (b1k+b2)%k) % k  = (a2b2)%k
// a^n % k = (a % k)^n % k


func myPow2(a int,k int,base int)  int {
	if k == 0 {
		return 1
	}

	if k == 1 {
		return a % base
	}

	a %= base
	if (k & 1) == 1 {
		return (a * myPow2(a,k-1,base)) % base
	}else{
		v :=  myPow2(a, k >> 1 ,base)
		return (v *v ) % base
	}
}

//a[1564] = a^4 * (a^156)^10
func superPowHelp(a int, b []int,pointer int) int {
	if a == 0 {
		return 0
	}

	if pointer == -1  {
		return 1
	}

	last := b[pointer]
	pointer = pointer -1

	p1 := myPow2(a,last,1337)
	p2 := myPow2(superPowHelp(a,b,pointer),10,1337)

	return (p1 * p2) % 1337


}

func superPower(a int,b []int) int  {
	return superPowHelp(a,b,len(b)-1)

}


func main() {
	fmt.Println(superPower(2,[]int{1,2}))
}
