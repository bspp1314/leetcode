package main

//Hash
func twoSum(nums []int, target int) []int {
	l := len(nums)
	result := make([]int, 2, 2)
	tmp := make(map[int]int)

	for i := 0; i < l; i++ {
		if j, ok := tmp[target-nums[i]]; ok {
			result[0] = j + 1
			result[1] = i + 1
			return result
		} else {
			tmp[nums[i]] = i
		}
	}
	return result
}

//双指针法
func twoSum2(nums []int, target int) []int {
	p1 := 0
	p2 := len(nums) - 1
	result := make([]int, 2, 2)

	for p1 != p2 {
		p := nums[p1] + nums[p2]
		if p > target {
			p2--
		} else if p < target {
			p1++
		} else {
			result[0] = p1 + 1
			result[1] = p2 + 1
			return result
		}
	}

	return result
}

func twoSum3(numbers []int, target int) []int {
	if len(numbers) <= 1 {
		return []int{}
	}


	left := 0
	right := len(numbers) -1
	for left < right {
		sum := numbers[left] + numbers[right]
		if sum == target {
			return []int{left+1,right+1}
		}else if sum > target{
			for left < right && numbers[right] == numbers[right-1] {
				right--
			}
			right--
		}else{
			for left < right && numbers[left] == numbers[left+1] {
				left++
			}
			left++
		}
	}

	return []int{}
}

func main() {

}
