package main

import "fmt"
import "math"

func reverse1(val int) int {
	var res int64
	for val != 0 {
		res = res*10 + int64(val%10)
		if res > math.MaxInt32 || res < math.MinInt32 {
			return 0
		}
		val = val / 10
	}
	return int(res)
}

func reverse2(x int) int {
	if x == 0 {
		return 0
	}

	val  := int32(x)

	var res int32

	for val != 0 {
		tem := res
		res = res*10 + val%10
		if tem * res < 0 {
			return 0
		}

		val = val / 10
	}

	return int(res)

}


func main() {
	fmt.Println(reverse1(math.MaxInt32))
	fmt.Println(reverse2(-2147483412))
}
