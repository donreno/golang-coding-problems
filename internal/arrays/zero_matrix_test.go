package arrays_test

import (
	"testing"

	goblin "github.com/franela/goblin"

	a "golang-coding-problems/internal/arrays"
)

func TestZeroMatrix(t *testing.T) {

	g := goblin.Goblin(t)

	g.Describe("Zero Matrix", func() {
		var m [][]int

		g.BeforeEach(func() {
			m = [][]int{}
		})

		g.It("Should zero matrix correctly", func() {
			m = [][]int{
				{1, 0, 1},
				{1, 1, 1},
				{1, 1, 0},
			}

			a.ZeroMatrix(m)

			g.Assert(m).Eql([][]int{
				{0, 0, 0},
				{1, 0, 0},
				{0, 0, 0},
			})
		})

		g.It("Should leave matrix unchanged if input is empty", func() {
			a.ZeroMatrix(m)

			g.Assert(m).Eql([][]int{})
		})
	})
}
