package main

import "fmt"

type Node struct {
	Left  int
	Right int
	Pre   *Node
	Next  *Node
}

func (n1 *Node) InsertPre(n2 *Node) {
	if n1.Pre == nil {
		n1.Pre = n2
		n2.Next = n1
	} else {
		n1.Pre.Next = n2
		n2.Pre = n1.Pre
		n1.Pre = n2
		n2.Next = n1
	}
}

func (n1 *Node) InsertNext(n2 *Node) {
	if n1.Next == nil {
		n1.Next = n2
		n2.Pre = n1
	} else {
		n1.Next.Pre = n2
		n2.Next = n1.Next
		n1.Next = n2
		n2.Pre = n1
	}
}

type RangeModule struct {
	intervals [][]int
}

func Constructor() RangeModule {
	return RangeModule{
		intervals: make([][]int, 0),
	}
}

func (this *RangeModule) AddRange(left int, right int) {

	res := make([][]int, 0)
	merged := false

	for _, interval := range this.intervals {
		// [left,right） 和  interval 没有交集,且在[left,right） 的在 interval 的右侧）
		if interval[0] > right {
			if !merged {
				res = append(res, []int{left, right})
				merged = true
			}

			res = append(res, interval)
		} else if interval[1] < left {
			// [left,right） 和  interval 没有交集,且在[left,right） 的在 interval 的左侧
			res = append(res, interval)
		} else {
			// 与插入区间有交集，计算它们的并集
			left = Min(left, interval[0])
			right = Max(right, interval[1])
		}

	}

	if !merged {
		res = append(res, []int{left, right})
	}

	this.intervals = res
}

func Min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func Max(a, b int) int {
	if a > b {
		return a
	}

	return b
}

func (this *RangeModule) QueryRange(left int, right int) bool {
	for _, interval := range this.intervals {
		if left >= interval[0] && right <= interval[1] {
			return true
		}
	}

	return false
}

func (this *RangeModule) RemoveRange(left int, right int) {
	res := make([][]int, 0)
	for _, interval := range this.intervals {
		// [left,right） 和  interval 没有交集,且在[left,right） 的在 interval 的右侧）
		// [left,right） 和  interval 没有交集,且在[left,right） 的在 interval 的左侧
		if interval[0] > right || interval[1] < left {
			res = append(res, interval)
		} else {
			if left <= interval[0] && right >= interval[1] {
				continue
			} else if left > interval[0] && right < interval[1] {
				//分裂成两个点
				res = append(res, []int{interval[0], left})
				res = append(res, []int{right, interval[1]})
			} else if left > interval[0] && right >= interval[1] {
				res = append(res, []int{interval[0], left})
			} else {
				res = append(res, []int{right, interval[1]})
			}
		}
	}

	this.intervals = res

}

func main() {
	// addRange(10, 20): null
	//removeRange(14, 16): null
	//queryRange(10, 14): true （区间 [10, 14) 中的每个数都正在被跟踪）
	//queryRange(13, 15): false （未跟踪区间 [13, 15) 中像 14, 14.03, 14.17 这样的数字）
	//queryRange(16, 17): true （尽管执行了删除操作，区间 [16, 17) 中的数字 16 仍然会被跟踪）
	r := Constructor()
	r.AddRange(6, 8)
	// [6,8)
	r.RemoveRange(7, 8)
	//[6,7)
	r.RemoveRange(8, 9)
	//[6,7)
	r.AddRange(8, 9)
	//[6,7) [8,9)
	r.RemoveRange(1, 3)
	//[6,7) [8,9)

	fmt.Println(r.intervals)
	r.AddRange(1, 8)
	fmt.Println(r.intervals)

}
