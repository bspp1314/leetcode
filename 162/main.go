package main


//你可以假设 nums[-1] = nums[n] = -∞ 。
func findPeakElement(nums []int) int {
	for i:=0;i < len(nums) -1;i++{
		if nums[i] > nums[i+1] {
			return i
		}
	}

	return len(nums)  - 1
}

func findPeakElement2(nums []int) int   {
	left := 0
	right := len(nums) -1

	for left <= right {
		mid := left + (right-left) >> 1

		if nums[mid] > nums[mid + 1] {
			right = mid
		} else {
			left = mid + 1
		}
	}

	return  left
}

func main() {

}
