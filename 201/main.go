package main

import "fmt"

//[5, 7]里共有三个数字，分别写出它们的二进制为：

//101　　110　　111

//相与后的结果为100，仔细观察我们可以得出，最后的数是该数字范围内所有的数的左边共同的部分，如果上面那个例子不太明显，我们再来看一个范围[26, 30]，它们的二进制如下：

// 11010　　11011　　11100　　11101　　11110


func rangeBitwiseAnd(m int, n int) int {
	i := 0
	for m != n {
		m = m >> 1
		n = n >> 1
		i++
	}

	return m << i
}

func main() {
	fmt.Println(rangeBitwiseAnd(17,20) )

}
