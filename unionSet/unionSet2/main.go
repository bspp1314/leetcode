package main

import (
	"errors"
	"fmt"
)

type unionSet struct {
	Parent []int
}

func NewUnionSet(size int) *unionSet {
	buf := make([]int, size)
	for i := 0; i < size; i++ {
		buf[i] = i  // 初始时，所有节点均指向自己
	}
	return &unionSet{Parent: buf}
}

func (s *unionSet) GetRoot(p int) (int, error) {
	if p < 0 || p >  len(s.Parent) {
		return 0, errors.New(
			"failed to get ID,index is illegal. ")
	}
	return s.getRoot(p), nil
}

func (s *unionSet) getRoot(p int) int {
    if s.Parent[p] == p {
    	return p
	}

	return  s.getRoot(s.Parent[p])
}

func (s *unionSet) IsConnected(p, q int) (bool, error) {
	if p < 0 || p > len(s.Parent) || q < 0 || q > len(s.Parent) {
		return false, errors.New(
			"error: index is illegal. ")
	}

	return s.getRoot(s.Parent[p]) == s.getRoot(s.Parent[p]), nil
}

func (s *unionSet)Union(p,q int) error  {
	if p < 0 || p > len(s.Parent) || q < 0 || q > len(s.Parent) {
		return errors.New(
			"error: index is illegal. ")
	}

	pRoot := s.getRoot(p)
	qRoot := s.getRoot(q)
	if pRoot != qRoot {
		s.Parent[pRoot] = qRoot
	}

	return nil
}




func main() {
	s := NewUnionSet(10)

	s.Union(0,1)
	s.Union(0,2)
	s.Union(0,3)
	s.Union(0,4)
	fmt.Println(s.Parent)
	fmt.Println(s.GetRoot(0))
}
