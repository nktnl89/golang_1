package search

import "testing"

func BenchmarkMyFindSimpleNumbers(b *testing.B) {
	search := MySimpleSearch{}
	seq := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15}
	for i := 0; i < b.N; i++ {
		search.FindSimpleNumbers(seq)
	}
}

func BenchmarkBigIntFindSimpleNumbers(b *testing.B) {
	search := StandardSearch{}
	seq := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15}
	for i := 0; i < b.N; i++ {
		search.FindSimpleNumbers(seq)
	}
}
