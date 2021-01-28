package sorting_test

import (
	s "golang-coding-problems/internal/sorting"
	"testing"

	goblin "github.com/franela/goblin"
)

func TestSearchInRotatedArray(t *testing.T) {
	g := goblin.Goblin(t)

	g.Describe("Search in rotated array", func() {

		arr := []int{15, 16, 19, 20, 25, 1, 3, 4, 5, 7, 10, 14}

		g.It("Should find correct index of elemenet", func() {
			g.Assert(s.SearchInRotatedArray(arr, 15)).Eql(0)
			g.Assert(s.SearchInRotatedArray(arr, 16)).Eql(1)
			g.Assert(s.SearchInRotatedArray(arr, 19)).Eql(2)
			g.Assert(s.SearchInRotatedArray(arr, 20)).Eql(3)
			g.Assert(s.SearchInRotatedArray(arr, 25)).Eql(4)
			g.Assert(s.SearchInRotatedArray(arr, 1)).Eql(5)
			g.Assert(s.SearchInRotatedArray(arr, 3)).Eql(6)
			g.Assert(s.SearchInRotatedArray(arr, 4)).Eql(7)
			g.Assert(s.SearchInRotatedArray(arr, 5)).Eql(8)
			g.Assert(s.SearchInRotatedArray(arr, 7)).Eql(9)
			g.Assert(s.SearchInRotatedArray(arr, 10)).Eql(10)
			g.Assert(s.SearchInRotatedArray(arr, 14)).Eql(11)
		})

		g.It("Should return -1 if element is not on array", func() {
			g.Assert(s.SearchInRotatedArray(arr, 11)).Eql(-1)
		})
	})
}
