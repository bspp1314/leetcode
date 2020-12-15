package unionSet2

import "testing"

func BenchmarkUnionSet_Union1(b *testing.B) {

	for j := 0; j < b.N; j++ {
		count := 100000
		s := NewUnionSet(count)
		for i := 1; i < count; i++ {
			s.Find(i)
			s.Union(i-1, i)
		}
	}

}

func BenchmarkUnionSet_Union2(b *testing.B) {

	for j := 0; j < b.N; j++ {
		count := 100000
		s := NewUnionSet(count)
		for i := 1; i < count; i++ {
			s.Find(i)
			s.Union(0, i)
		}
	}
}


