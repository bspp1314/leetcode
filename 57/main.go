package main

func Min(a,b int) int   {
	if a < b {
		return a
	}
	return b
}

func Max(a,b int) int  {
	if a > b {
		return a
	}

	return b
}
func insert(intervals [][]int, newInterval []int) (ans [][]int) {
	left, right := newInterval[0], newInterval[1]
	merged := false

	for _, interval := range intervals {
		if interval[0] > right {
			if !merged {
				ans = append(ans, []int{left, right})
				merged = true
			}
			ans = append(ans,interval)
		}else if interval[1] < left {
			// 在插入区间的左侧且无交集
			ans = append(ans, interval)
		}else{
			// 与插入区间有交集，计算它们的并集
			left = Min(left, interval[0])
			right = Max(right, interval[1])

		}
	}

	if !merged {
		ans = append(ans, []int{left, right})
	}
	return
}

func main() {

}
