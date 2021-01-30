package sorting_test

import (
	s "golang-coding-problems/internal/sorting"
	"testing"

	goblin "github.com/franela/goblin"
)

func TestSortedMatrixSearch(t *testing.T) {
	g := goblin.Goblin(t)

	g.Describe("Sorter matrix search", func() {

		var m [][]int

		g.It("Should find position for element present on matrix", func() {
			r1, c1 := s.FindElementInSortedMatrix(m, 55)
			r2, c2 := s.FindElementInSortedMatrix(m, 85)

			g.Assert(r1).Eql(2)
			g.Assert(c1).Eql(1)
			g.Assert(r2).Eql(0)
			g.Assert(c2).Eql(3)
		})

		g.It("Should not find element that's not on matrix", func() {
			r, c := s.FindElementInSortedMatrix(m, 21)

			g.Assert(r).Eql(-1)
			g.Assert(c).Eql(-1)
		})

		g.BeforeEach(func() {
			m = [][]int{
				{15, 20, 70, 85},
				{20, 35, 80, 95},
				{30, 55, 95, 105},
				{40, 80, 100, 120},
			}
		})

	})
}
