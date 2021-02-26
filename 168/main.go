package main

import "fmt"

func convertIntToByte(i int) byte {
	if  i == 1 {
		return 'A'
	}


	return byte(int(convertIntToByte(1)) + (i - 1))
}
func convertToTitle(n int) string {
	res := make([]byte,0)

	for n != 0 {
		n--
		res = append(res,byte(n % 26 + 'A'))
		n = n / 26
	}

	left := 0
	right := len(res) -1

	for left   < right{
		 res[left],res[right] = res[right],res[left]
		 left++
		 right--
	}

	return string(res)
}

func main() {
	fmt.Println(convertToTitle(27))
}
