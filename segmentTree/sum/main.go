package sum

type SegmentTree struct {
	SumTree []int //线段树
	Data    []int //数组数据
}

func Max(a,b int) int   {
	if a > b {
		return a
	}

	return b
}

func Min(a,b int) int   {
	if a < b {
		return a
	}

	return b
}

func NewSegmentTree(arr []int) *SegmentTree {
	n := len(arr)
	m := 0
	if (n&n - 1) == 0 {
		m = 2 * n
	} else {
		m = 4 * n
	}
	s := &SegmentTree{
		SumTree: make([]int, m),
		Data:    arr,
	}

	return s
}

// 构建以index为根的树
func (s *SegmentTree) build(index int, left, right int) {
	if left == right {
		s.SumTree[index] = left
	} else {
		mid := left + (right-left)/2
		leftIndex := index << 1
		rightIndex := leftIndex + 1
		s.build(leftIndex, left, mid)
		s.build(rightIndex, mid+1, right)
		s.SumTree[index] = s.SumTree[leftIndex] + s.SumTree[rightIndex]
	}
}

