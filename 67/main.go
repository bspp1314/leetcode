package main

import "fmt"

func ByteToInt(b byte) int  {
	return int(b - '0')
}

func IntToByte(i int) byte  {
		return byte(i) + '0'
}
//leetcode submit region begin(Prohibit modification and deletion)
func addBinary(a string, b string) string {
	if a == "0" || a == "" {
		return b
	}

	if b == "0" || b == "" {
		return a
	}

	aLen := len(a)
	bLen := len(b)

	i := aLen -1
	j := bLen -1

	var res []byte
	c := 0
	for i >= 0 || j >= 0 {
		if i >= 0 {
			c += ByteToInt(a[i])
			i--
		}

		if j >= 0 {
			c += ByteToInt(b[j])
			j--
		}

		res = append(res,IntToByte(c & 1))
		c = c >> 1
	}

	if c > 0 {
		res = append(res,IntToByte(c))
	}

	left := 0
	right := len(res) -1

	for left < right {
		res[left],res[right] = res[right],res[left]
		left++
		right--
	}

	return string(res)
}


func main() {
	fmt.Printf("=====%s=====\n",addBinary("1111", "1111"))
	//fmt.Printf("byte to int %d \n",ByteToInt('9'))
	//fmt.Printf("inte to byte %c \n",IntToByte(9))
}
