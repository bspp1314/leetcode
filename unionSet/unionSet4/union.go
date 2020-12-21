package unionSet4

import "fmt"

type unionSet struct {
	rank []int // rank[i]表示以i为根的树的高度
	set  []int
}

func NewUnionSet(size int) *unionSet {
	u := &unionSet{
		rank: make([]int, size),
		set:  make([]int, size),
	}
	for i := 0; i < size; i++ {
		u.rank[i] = 1
		u.set[i] = i
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
	for i != s.set[i] {
		// i->parent = i->parent->parent
		s.set[i] = s.set[s.set[i]]
		i = s.set[i]
	}

	return i
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

	iRoot := s.getRoot(i)
	jRoot := s.getRoot(j)
	if iRoot != jRoot {
		if s.rank[iRoot] < s.rank[jRoot] {
			s.set[iRoot] = jRoot
		}else if s.rank[jRoot] > s.rank[iRoot]{
			s.set[jRoot] = iRoot
		}else{
			s.set[iRoot] = jRoot
			s.rank[jRoot] += 1
		}
	}

	return nil
}

//func main() {
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
//
//
//
