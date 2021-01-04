package arrays_test

import (
	"testing"

	goblin "github.com/franela/goblin"

	a "golang-coding-problems/internal/arrays"
)

func TestRotateMatrixV1(t *testing.T) {
	g := goblin.Goblin(t)

	g.Describe("Rotate Matrix 90 Degrees Clockwise", func() {
		var m [][]int

		g.BeforeEach(func() {
			m = [][]int{
				{1, 2, 3},
				{4, 5, 6},
				{7, 8, 9},
			}
		})

		g.It("Should rotate correctly", func() {
			a.RotateMatrix90Deg(m)

			g.Assert(m).Eql([][]int{
				{7, 4, 1},
				{8, 5, 2},
				{9, 6, 3},
			})
		})

		g.It("Should not rotate if matrix has just 1 row", func() {
			m = [][]int{{1, 2, 3}}

			a.RotateMatrix90Deg(m)

			g.Assert(m).Eql([][]int{{1, 2, 3}})
		})

		g.It("Should not rotate if matrix is empty", func() {
			m = [][]int{}

			a.RotateMatrix90Deg(m)

			g.Assert(m).Eql([][]int{})
		})
	})
}
