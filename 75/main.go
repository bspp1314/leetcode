package main

func sortColors(nums []int)  {
	slow := 0
	fast := 0
	right := len(nums) -1

	for fast <= right {
		if nums[fast] == 0 {
			nums[fast],nums[slow] = nums[slow],nums[fast]
			slow++
			fast++
		}else if nums[fast] == 2 {
			nums[fast],nums[right] = nums[right],nums[fast]
			right--
		}else{
			fast++
		}
	}


}
// 2 2 2 1 1 1 0 0 0
func main()  {
	sortColors([]int{0,0,0,1,1,1,2,2,2})
}
