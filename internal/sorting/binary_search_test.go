package sorting_test

import (
	s "golang-coding-problems/internal/sorting"
	"testing"

	goblin "github.com/franela/goblin"
)

func TestBinarySearch(t *testing.T) {
	g := goblin.Goblin(t)

	g.Describe("Binary Search", func() {

		g.It("Should return correct index for found element", func() {
			g.Assert(s.BinarySearch([]int{1, 2, 3}, 2)).Eql(1)
			g.Assert(s.BinarySearch([]int{12, 36, 90}, 12)).Eql(0)
			g.Assert(s.BinarySearch([]int{12, 36, 90}, 150)).Eql(-1)
		})

		g.It("Should return correct index for found element (Recursive)", func() {
			g.Assert(s.BinarySearchRecursive([]int{1, 2, 3}, 2)).Eql(1)
			g.Assert(s.BinarySearchRecursive([]int{12, 36, 90}, 12)).Eql(0)
			g.Assert(s.BinarySearchRecursive([]int{12, 36, 90}, 150)).Eql(-1)
		})
	})
}
