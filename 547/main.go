package main

import (
	"fmt"
)

// 实现
type Set struct {
	rank []int // rank[i]表示以i为根的树的高度
	Set  []int
}

func NewUnionSet(size int) *Set {
	s := &Set{
		rank: make([]int,size),
		Set:  make([]int,size),
	}
	for i := 0; i < size; i++ {
		s.rank[i] = 1
		s.Set[i] = i
	}

	return s
}


func (set *Set) getRoot(p int) int {
	for p != set.Set[p] {
		set.Set[p] = set.Set[set.Set[p]]
		p = set.Set[p]
	}
	return p
}


func (set *Set) Union(p, q int) error {

	pRoot := set.getRoot(p)
	qRoot := set.getRoot(q)


	if pRoot != qRoot {
		if set.rank[pRoot] < set.rank[qRoot] {
			set.Set[pRoot] = qRoot
		} else if set.rank[qRoot] < set.rank[pRoot] {
			set.Set[qRoot] = pRoot
		} else {
			set.Set[pRoot] = qRoot
			set.rank[qRoot] += 1
		}
	}
	return nil
}
func (s *Set) Num() int  {
	res := 0
	for i := 0; i < len(s.Set); i++ {
		if s.Set[i] == i {
			res ++
		}
	}

	return res
}

func findCircleNum2(M [][]int) int   {
	s := NewUnionSet(len(M))

	fmt.Println(s.Set)
	for i := 0; i < len(M); i++ {
		for j := 0; j < len(M[i]); j++ {
			if M[i][j] == 1 && i != j {
				s.Union(i,j)
			}
		}
	}

	return s.Num()
}

func findCircleNum(M [][]int) int {
	visited := make([]bool, len(M))

	queue := make([]int, 0)
	res := 0

	for i := 0; i < len(M); i++ {
		if visited[i] {
			continue
		}
		queue = append(queue, i)

		for len(queue) != 0 {
			s := queue[0]
			queue = queue[1:]

			// 访问点s相邻的节点
			for j := 0; j < len(M); j++ {
				if M[s][j] == 1 && !visited[j] {
					queue = append(queue, j)
				}
			}
		}

		res++
	}

	return res
}

func main() {
	out := findCircleNum2([][]int{
		[]int{1,1,0},
		[]int{1,1,0},
		[]int{0,0,1},
	})

	fmt.Println(out)

	out2 := findCircleNum2([][]int{
		[]int{1,0,0,1},
		[]int{0,1,1,0},
		[]int{0,1,1,1},
		[]int{1,0,1,1},
	})

	fmt.Println(out2)

}
