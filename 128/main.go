package main

func longestConsecutive(nums []int) int {
	if len(nums) <= 1 {
		return len(nums)
	}

	numsSet := make(map[int]bool)

	for _, num := range nums {
		numsSet[num] = true
	}

	res := 0

	for num, _ := range numsSet {
		if !numsSet[num-1] {
			beginNum := num+1
			newRes  := 1

			for numsSet[beginNum] {
				beginNum++
				newRes++
			}

			if res < newRes {
				res = newRes
			}
		}
	}

	return res

}




func main() {

}
