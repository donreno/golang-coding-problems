package sorting_test

import (
	s "golang-coding-problems/internal/sorting"
	"testing"

	goblin "github.com/franela/goblin"
)

func TestSearchSortedNoSize(t *testing.T) {
	g := goblin.Goblin(t)

	g.Describe("Search Sorted No Size (Listy)", func() {

		var l s.Listy

		g.BeforeEach(func() {
			l = s.Listy([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10})
		})

		g.It("Should find expected index for element on Listy", func() {

			g.Assert(s.SortedSearchNoSize(l, 1)).Eql(0)
			g.Assert(s.SortedSearchNoSize(l, 5)).Eql(4)
			g.Assert(s.SortedSearchNoSize(l, 9)).Eql(8)
			g.Assert(s.SortedSearchNoSize(l, 10)).Eql(9)
			g.Assert(s.SortedSearchNoSize(l, 20)).Eql(-1)
		})
	})
}
