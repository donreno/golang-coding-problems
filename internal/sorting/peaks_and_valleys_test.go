package sorting_test

import (
	s "golang-coding-problems/internal/sorting"
	"testing"

	goblin "github.com/franela/goblin"
)

func TestPeaksAndValleys(t *testing.T) {
	g := goblin.Goblin(t)

	g.Describe("Peaks and Valleys", func() {
		var input []int

		g.It("Should build expected peaks and valleys arr", func() {
			s.PeaksAndValleys(input)
			g.Assert(input).Eql([]int{3, 5, 1, 2, 3})
		})

		g.BeforeEach(func() {
			input = []int{5, 3, 1, 2, 3}
		})
	})
}
