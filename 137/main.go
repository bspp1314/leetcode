package main

import "fmt"

// 一个二进制位只能表示0或者1。也就是天生可以记录一个数出现了一次还是两次。
// x ^ 0 = x
// x ^ x = 0

// 要记录出现3次，需要两个二进制位。那么上面单独的x就不行了。我们需要两个变量，每个变量取一位

// ab ^ 00 = ab
// ab ^ ab = 00

//b = x &


//x = x[7] x[6] x[5] x[4] x[3] x[2] x[1] x[0]
//x = (a[7]b[7]) (a[6]b[6]) ... (a[1]b[1]) (a[0]b[0])

//它是一个逻辑电路，a、b变量中，相同位置上，分别取出一位，负责完成00->01->10->00，也就是开头的那句话，当数字出现3次时置零。
func singleNumber(nums []int) int {
	a := 0
	b := 0
	for _,x := range nums {
		b = (b ^ x) & (^a)
		a = (a ^ x) & (^b)

		fmt.Println("a",a)
		fmt.Println("b",b)
	}

	return b

}

func main() {
	fmt.Println(singleNumber([]int{1, 1, 1}))

}
