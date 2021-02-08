package unionSet2

import (
	"fmt"
)

type unionSet struct {
	set []int
}

func NewUnionSet(size int) *unionSet {
	u  := &unionSet{set: make([]int,size)}
	for i := 0; i < size; i++ {
		u.set[i] = i  // 初始时，所有节点均指向自己
	}

	return u
}

func (s *unionSet)GetSet() []int  {
	return s.set
}


func (s *unionSet) Find(i int) (int, error) {
	if i < 0 || i >  len(s.set) {
		return 0, fmt.Errorf(
			"Index is illegal. ")
	}
	return s.getRoot(i), nil
}

func (s *unionSet) getRoot(i int) int {
    if s.set[i] == i {
    	return i
	}

	return  s.getRoot(s.set[i])
}

func (s *unionSet) IsConnected(i, j int) (bool, error) {
	if i < 0 || i > len(s.set) || j < 0 || j > len(s.set) {
		return false, fmt.Errorf(
			"Index is illegal. ")
	}

	return s.getRoot(s.set[i]) == s.getRoot(s.set[j]), nil
}

func (s *unionSet)Union(i,j int) error  {
	if i < 0 || i > len(s.set) || j < 0 || j > len(s.set) {
		return fmt.Errorf(
			"Index is illegal. ")
	}

	pRoot := s.getRoot(i)
	qRoot := s.getRoot(j)
	if pRoot != qRoot {
		s.set[pRoot] = qRoot
	}

	return nil
}




//func 492() {
//	s := NewUnionSet(10)
//	//Before union
//	fmt.Println(s.GetSet())
//
//	s.Union(0,1)
//	s.Union(0,2)
//	s.Union(0,3)
//
//	fmt.Println(s.GetSet())
//
//	s.Union(4,5)
//	s.Union(4,6)
//	s.Union(4,7)
//
//	fmt.Println(s.GetSet())
//
//
//	s.Union(2,4)
//
//	fmt.Println(s.GetSet())
//}
