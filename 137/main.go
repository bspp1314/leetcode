package main

import "fmt"

// (0,0) 0 (0,0)
// (0,1) 0 (0,1)
// (1,0) 0 (1,0)
// (0,0) 1 (0,0)
// (0,1) 1 (0,1)
// (1,0) 1 (1,0)

// newA = a & ~b & ~n | ~a & b & n
// newB = ~a & b & ~n | ~a & ~b & n

func singleNumber(nums []int) int {
	a := 0
	b := 0
	for n := range nums {
		newA := a & ^b & ^n | ^a & b & n
		newB := ^a & b & ^n | ^a & ^b & n

		newA,newB = newB,newA
	}

	return b

}

func main() {
	fmt.Println(singleNumber([]int{1, 1, 1, 3}))
	//fmt.Printf("(%d,%d) ==> (%d,%d)\n",a,b,a & ^b  ^ 0,^a & b  ^ 0)
	//a = 0
	//b = 1
	//fmt.Printf("(%d,%d) ==> (%d,%d)\n",a,b,a & ^b  ^ 0,^a & b  ^ 0)
	//a = 1
	//b = 0
	//fmt.Printf("(%d,%d) ==> (%d,%d)\n",a,b,a & ^b  ^ 0,^a & b  ^ 0)
	//

	//a = 0
	//b = 0
	//fmt.Printf("(%d,%d) ==> (%d,%d)\n",a,b,a & ^b  ^ ^1,10)
	//a = 0
	//b = 1
	//fmt.Printf("(%d,%d) ==> (%d,%d)\n",a,b,a & ^b  ^ ^1,10)
	//a = 1
	//b = 0
	//fmt.Printf("(%d,%d) ==> (%d,%d)\n",a,b,a & ^b  ^ ^1,10)

}
