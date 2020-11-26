package main

import "fmt"


func reverseBits2(num uint32) uint32 {
	res := uint32(0)

	for i:=0;i<=31;i++ {
		res = (res << 1) |  (num & 1)
		num = num >> 1
	}

	return res
}


func reverseBits(num uint32) uint32 {
	num = (num >> 16) | (num << 16 )
	// 11111111000000001111111100000000   00000000111111110000000011111111
	num = (num & 0xff00ff00) >> 8 |  ((num & 0x00ff00ff) << 8)
	// 11110000111100001111000011110000   00001111000011110000111100001111
	num = (num & 0xf0f0f0f0)  >> 4 | ((num & 0x0f0f0f0f) << 4)
	// 11001100110011001100110011001100   00110011001100110011001100110011
	num = (num & 0xcccccccc) >> 2 |  ((num & 0x33333333) << 2)
	// 10101010101010101010101010101010   010101010101010101010101010101010
	num = (num & 0xaaaaaaaa) >> 1 |  ((num & 0x55555555) << 1)
	return num
}

func main() {
	out := uint32(1)
	fmt.Printf("%032b\n",out)
	fmt.Printf("%032b\n",reverseBits(out))
	fmt.Printf("%032b\n",reverseBits2(out))

	//
	// 去掉最后一位  x >> 1
	// 在最后加一个0 x << 1
	// 在最后一位加1 x << 1  + 1
	// 把最后一位变成1 x | 1
	// 把最后一位变成0 x | 1 + 1
	// 最后一位取反    x ^ 1
	// 把右数第k位变成1  (101001->101101,k=3)  x | (1 << (k -1 ))
	// 把右数第k位变成0  (101101->101001,k=3)  x & (^(1 << (k -1 )))
}
