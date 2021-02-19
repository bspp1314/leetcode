package main

import "fmt"

func insert2(intervals [][]int, newInterval []int) (res [][]int)  {
	left,right := newInterval[0],newInterval[1]
	merged := false

	for _, interval := range intervals {
		// （left,right） 和  interval 没有交集,且在（left,right 的在 interval 的右侧）
		if interval[0] > right {
			if ! merged {
			   res = append(res,[]int{left,right})
			   merged = true
			}

			res = append(res,interval)
		}else if interval[1] < left {
			// left,right） 和  interval 没有交集,且在（left,right 的在 interval 的左侧
			res = append(res,interval)
		}else{
			// 与插入区间有交集，计算它们的并集
			left = Min(left, interval[0])
			right = Max(right, interval[1])
		}
	}

	if !merged {
		res  = append(res,[]int{left,right})
	}

	return res
}

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
func insert(intervals [][]int, newInterval []int) [][]int {
	if len(intervals) == 0 {
		return [][]int{newInterval}
	}

	temp := make([][]int, len(intervals)+1)
	if intervals[0][0] >= newInterval[0] {
		temp[0] = newInterval
		copy(temp[1:], intervals)
	} else if intervals[len(intervals)-1][0] <= newInterval[0] {
		temp[len(temp)-1] = newInterval
		copy(temp[:len(temp)-1], intervals)
	} else {
		for i := 0; i < len(intervals); i++ {
			if intervals[i][0] >= newInterval[0] {
				copy(temp[0:i], intervals[0:i])
				temp[i] = newInterval
				copy(temp[i+1:len(temp)], intervals[i:len(intervals)])
				break
			}
		}
	}

	res := [][]int{temp[0]}
	for i := 1; i < len(temp); i++ {
		if temp[i][0] <= res[len(res)-1][1] {
			//merge
			if temp[i][1] > res[len(res)-1][1] {
				res[len(res)-1][1] = temp[i][1]
			}
		} else {
			res = append(res, temp[i])
		}
	}

	return res
}

func main() {
	//输入：intervals = [[1,3],[6,9]], newInterval = [2,5]
	//输出：[[1,5],[6,9]]
	fmt.Println(insert([][]int{
		{1, 3},
		{6, 9},
	}, []int{2, 5}))

	//输入：intervals = [[1,2],[3,5],[6,7],[8,10],[12,16]], newInterval = [4,8]
	//输出：[[1,2],[3,10],[12,16]]
	fmt.Println(insert([][]int{
		{1,2},
		{3,5},
		{6,7},
		{8,10},
		{12,16},
	},[]int{4,8}))
	//
	////输入：intervals = [], newInterval = [5,7]
	////输出：[[5,7]]
	fmt.Println(insert([][]int{
	},[]int{5,7}))
	//
	////输入：intervals = [[1,5]], newInterval = [2,3]
	////输出：[[1,5]]
	fmt.Println(insert([][]int{
		{1,5},
	},[]int{2,3}))
	//
	//
	////输入：intervals = [[1,5]], newInterval = [2,7]
	////输出：[[1,7]]
	fmt.Println(insert([][]int{
		{1,5},
	},[]int{2,7}))

}
