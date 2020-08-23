package main

func sortColors(nums []int)  {

	leftSlow := 0
	leftFast := 0
	right := len(nums) -1

	leftSlow = leftFast

	for leftFast <= right {
		if nums[leftFast] == 2 {
			nums[leftFast],nums[right] = nums[right],nums[leftFast]
			right --
		}else if nums[leftFast] == 1 {
			leftFast++
		}else {
			nums[leftFast],nums[leftSlow] = nums[leftSlow],nums[leftFast]
			leftFast++
			leftSlow++
		}
	}
}

func main()  {
	sortColors([]int{0,0,0,1,1,1,2,2,2})
}
