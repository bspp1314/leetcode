package main

func max(i, j int) int {
	if i > j {
		return i
	}

	return j
}
func canJump(nums []int) bool {
	n := len(nums)
	rightmost := 0
	for i := 0; i < n; i++ {
		if i <= rightmost {
			rightmost = max(rightmost, i+nums[i])
			if rightmost >= n-1 {
				return true
			}
		}
	}

	return false

}

func main() {
}
