package sorting_test

import (
	s "golang-coding-problems/internal/sorting"
	"testing"

	goblin "github.com/franela/goblin"
)

func TestSortedMerge(t *testing.T) {
	g := goblin.Goblin(t)

	var a []int
	var b []int

	g.Describe("Sorted Merge", func() {

		g.It("Should fail if a is smaller or equal size than b", func() {
			a, b = []int{1}, []int{3, 4}
			err := s.SortedMerge(a, b)

			g.Assert(err).IsNotNil()
		})

		g.It("Should merge and sort arrays properly", func() {
			err := s.SortedMerge(a, b)

			g.Assert(err).IsNil()
			g.Assert(a).Eql([]int{0, 1, 2, 3, 4, 20})
		})

		g.BeforeEach(func() {
			a = make([]int, 6)
			a[0] = 20
			a[1] = 2
			a[2] = 3
			b = []int{0, 4, 1}
		})
	})
}
