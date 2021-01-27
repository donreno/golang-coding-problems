package sorting_test

import (
	s "golang-coding-problems/internal/sorting"
	"testing"

	goblin "github.com/franela/goblin"
)

func TestMergesort(t *testing.T) {
	g := goblin.Goblin(t)

	g.Describe("Merge Sort", func() {
		var arr []int

		g.BeforeEach(func() {
			arr = []int{9, 7, 2, 1, 4, 5, 0, 3, 6, 8}
		})

		g.It("Should sort array correctly", func() {
			s.MergeSort(arr)

			g.Assert(arr).Eql([]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9})
		})

	})
}

func BenchmarkMergeSort(b *testing.B) {
	for i := 0; i < b.N; i++ {
		s.MergeSort([]int{9, 7, 2, 1, 4, 5, 0, 3, 6, 8})
	}
}
