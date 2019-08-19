package main

import "fmt"
import "math"

func conver(val int) int {
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
func main() {
	fmt.Println(conver(123400))
	fmt.Println(conver(-123400))
}
