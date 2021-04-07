package main

import "math"

func splitIntoFibonacci(S string) []int {
	n := len(S)
	var back  func(index int,first,second int,count int,sub []int)([]int,bool)
	back = func(index int,first,second int, count int,sub []int) ([]int,bool) {
		if index == n {
			if count >= 3 {
				return sub,true
			}else{
				return []int{},false
			}
		}

		sum := first + second


		value := 0
		for i := index;i < n ;i++ {
			// 除 0 以外，其他数字第一位不能为 0
			if i > index && S[index] == '0' {
				break
			}

			value = value * 10 + int(S[i] - '0')

			if value > math.MaxInt32 {
				break
			}

			if count >= 2 {
				if value < sum {
					// 小的话继续向后继续拼接
					continue
				} else if value > sum {
					// 大的话直接结束，再往后拼接无意义
					break
				}
			}

			newSub := append(sub,value)

			if res,v := back(i + 1, second, value, count + 1,newSub);v{
				return res,v
			}
		}

		return []int{},false
	}

	v,_ := back(0,0,0,0,[]int{})
	return v
}

func main() {

}
