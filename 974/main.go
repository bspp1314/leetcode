package main

import "fmt"
func subarraysDivByK2(A []int, K int) int {
	record := map[int]int{0: 1}
	sum, ans := 0, 0
	for _, elem := range A {
		sum += elem
		modulus := (sum % K + K) % K
		ans += record[modulus]
		record[modulus]++
	}
	return ans
}

func subarraysDivByK(nums []int, k int) int {
	if len(nums) == 0 {
		return 0
	}
	var res int

	n := len(nums)

	for i := 0; i < n; i++ {
		sum := 0
		for j := i; j < n; j++ {
			if i == j {
				sum = nums[j]
			} else {
				sum += nums[j]
			}

			if sum%k == 0 {
				res++
			}
		}
	}

	return res

}

func main() {
	fmt.Println(subarraysDivByK([]int{4, 5, 0, -2, -3, 1}, 5))
}
