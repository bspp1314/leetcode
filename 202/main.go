package main

import "fmt"

//这题可以用快慢指针的思想去做，有点类似于检测是否为环形链表那道题
//如果给定的数字最后会一直循环重复，那么快的指针（值）一定会追上慢的指针（值），也就是
//两者一定会相等。如果没有循环重复，那么最后快慢指针也会相等，且都等于1。
func isHappy(n int) bool {
	slow := n
	fast := n

	for {
		slow = squareSum(slow)
		fast = squareSum(squareSum(fast))
		if slow == fast {
			break
		}
	}

	if fast == 1 {
		return true
	}else{
		return false
	}

}

var Repeat  = make(map[int]bool,0)
func isHappy2(n int) bool {
	if n == 1 {
		return true
	}

	if Repeat[n] {
		return false
	}

	Repeat[n] = true
	res := 0
	for n != 0 {
		k := n % 10
		res += k * k
		n = n / 10
	}

	return isHappy2(res)


}

func squareSum(n int) int   {
	res := 0
	for n != 0 {
		k := n % 10
		res += k * k
		n = n / 10
	}

	return res
}

func main() {
	fmt.Println(isHappy(2))
	fmt.Println(isHappy2(2))
}
