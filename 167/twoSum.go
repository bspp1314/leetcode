package main

func twoSum(nums []int,target int) []int  {
	if len(nums) <= 1 {
		return []int{}
	}
	left := 0
	right := len(nums) -1

	for left < right {
		sum := nums[left] + nums[right]
		if sum == target {
			return []int{left+1,right+1}
		}else if sum > target {
			right--
		}else{
			left++
		}
	}

	return []int{}
}


func main() {

}
