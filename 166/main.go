package main

import (
	"fmt"
)

func Abs(a int) int   {
	if a < 0 {
		return -a
	}

	return  a
}

func fractionToDecimal(numerator int, denominator int) string {
	if numerator == 0 {
		return "0"
	}

	res := make([]byte,0)

	if numerator ^ denominator  < 0 {
		res = append(res,'-')
	}

	numerator = Abs(numerator)
	denominator = Abs(denominator)

	res = append(res,[]byte(fmt.Sprintf("%d",numerator / denominator))...)


	numerator = numerator % denominator
	if numerator == 0 {
		return string(res)
	}

	res = append(res,'.')
	remainderMap := make(map[int]int)

	for numerator != 0  {
		if  index,exist := remainderMap[numerator];exist {
			newRes := make([]byte,0,len(res)+2)
			newRes = append(newRes,res[0:index]...)
			newRes = append(newRes,'(')
			newRes = append(newRes,res[index:]...)
			newRes = append(newRes,')')
			res = newRes
			break
		}

		remainderMap[numerator] = len(res)
		numerator *=10
		res = append(res,[]byte(fmt.Sprintf("%d",numerator / denominator))...)
		numerator = numerator % denominator
	}

	return string(res)
}

func main() {
	fmt.Println(fractionToDecimal(333,-333))
}
