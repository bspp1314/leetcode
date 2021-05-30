package main

import (
	"fmt"
	"strconv"
)

func ByteToInt(b byte) int {
	return int(b - '0')
}

func IntToByte(i int) byte {
	return byte(i + '0')
}


func addStrings(num1 string, num2 string) string {
	index := 0
	ans := ""
	for i, j := len(num1) - 1, len(num2) - 1; i >= 0 || j >= 0 || index != 0; i, j = i - 1, j - 1 {
		var x, y int
		if i >= 0 {
			x = int(num1[i] - '0')
		}
		if j >= 0 {
			y = int(num2[j] - '0')
		}
		result := x + y + index
		ans = strconv.Itoa(result%10) + ans
		index = result / 10
	}

	return ans
}

func multiply(num1 string, num2 string) string {
	if num1 == "0" || num2 == "0" {
		return "0"
	}
	m := len(num1)
	n := len(num2)
	results := make([]int, m+n)
	for i := m - 1; i >= 0; i-- {
		x := int(num1[i]) - '0'
		for j := n - 1; j >= 0; j-- {
			y := int(num2[j] - '0')
			results[i+j+1] += x * y
		}
	}

	for i := m + n - 1; i > 0; i-- {
		results[i-1] += results[i] / 10
		results[i] %= 10
	}


	result := ""
	idx := 0
	if results[0] == 0 {
		idx = 1
	}
	for ; idx < m+n; idx++ {
		result += strconv.Itoa(results[idx])
	}
	return result

}

//8111
// 333
func add(nums1,nums2 []byte) []byte {
	//if len(nums1) < len(nums2) {
	//	return add(nums2,nums1)
	//}

	p1 := len(nums1)-1
	p2 := len(nums2)-1
	low := 0
	res := make([]byte,0)
	for p1 >= 0 || p2 >= 0 {
		h := low
		if p1 >= 0 {
			h +=  int(nums1[p1] - '0')
			p1--
		}

		if p2 >= 0 {
			h +=  int(nums2[p2] - '0')
			p2--
		}

		res = append(res,byte('0' + h%10))
		low = h / 10
	}

	if low != 0 {
		res = append(res,byte('0'+low))
	}

	left := 0
	right := len(res) -1

	for left < right {
		res[right],res[left] = res[left],res[right]
		left++
		right--
	}

	return res
}

func multiply2(num1 string, num2 string) string {
	if num1 == "0" || num2 == "0" {
		return "0"
	}

	var res []byte
	for i := 0;i < len(num1);i ++ {
		if i == 0 {
			res = mul(num2,num1[i])
		}else{
			res = add(res,mul10(mul(num2,num1[i]),i))
		}
	}

	return string(res)
}

func mul10(num []byte,d int) []byte {
	if d == 0 {
		return num
	}


	for i := 0;i < d ;i++ {
		num = append(num,'0')
	}

	return num
}

func mul(num string,c byte) []byte {
	if c  == '0' {
		return []byte{'0'}
	}

	if c == '1' {
		return []byte(num)
	}

	p := len(num) -1
	low := 0
	res := make([]byte,0)

	for  p >= 0 {
		h := low
		h += int(num[p]-'0') * int(c - '0')
		res = append(res,byte(h+'0') % 10 )
		low = h / 10
		p--
	}

	if low != 0 {
		res = append(res,byte('0'+low))
	}

	left := 0
	right := len(res) -1

	for left < right {
		res[right],res[left] = res[left],res[right]
		left++
		right--
	}

	return res
}


func main() {
	fmt.Println(multiply2("3","245"))
}
