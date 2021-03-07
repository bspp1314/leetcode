package main

//个人理解摩尔投票法本质上就是宾果消消乐游戏，每次消除3个不同的数。由于数组长度为n，
//因此消消乐最多进行[n/3]次。因此，我们想要的答案（超过[n/3]的数字）一定没有被消除完，
//一定存在最后活下来的两个数当中。 但是，存活的两个数不一定都是想要的真正的答案，最后再遍历确认一下这两个数是不是答案即可。
func majorityElement(nums []int) []int {
	// 创建返回值
	var res = make([]int, 0)
	if nums == nil || len(nums) == 0 {
		return res
	}

	// 初始化两个候选人 candidate，以及他们的计数票
	value1 := nums[0]
	count1 := 0
	value2 := nums[0]
	count2 := 0

	//摩尔投票法
	// 配对阶段
	for _, num := range nums {
		// 投票
		if value1 == num {
			count1++
			continue
		}
		if value2 == num {
			count2++
			continue
		}

		if count1 == 0 {
			value1 = num
			count1++
			continue
		}
		if count2 == 0 {
			value2 = num
			count2++
			continue
		}

		count1--
		count2--
	}
	// 计数阶段
	count1 = 0
	count2 = 0
	for _, num := range nums {
		if value1 == num {
			count1++
		}else if value2 == num {
			count2++
		}
	}

	if count1 > len(nums)/3 {
		res = append(res, value1)
	}
	if count2 > len(nums)/3 {
		res = append(res, value2)
	}
	return res
}


func main() {

}
