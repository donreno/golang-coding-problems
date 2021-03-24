package crackingthecodinginterview_test

import (
	ctci "golang-coding-problems/internal/crackingthecodinginterview"

	"testing"

	"github.com/franela/goblin"
)

func TestGetMaxRegionSize(t *testing.T) {
	g := goblin.Goblin(t)

	g.Describe("Get Max Region Size", func() {
		var matrix [][]int

		g.It("Should get the expected max region", func() {
			g.Assert(ctci.GetMaxRegionSize(matrix)).Eql(7)
		})

		g.BeforeEach(func() {
			matrix = [][]int{
				{0, 0, 0, 1, 1, 0, 0},
				{0, 1, 0, 0, 1, 1, 0},
				{1, 1, 0, 1, 0, 0, 1},
				{0, 0, 0, 0, 0, 1, 0},
				{1, 1, 0, 0, 0, 0, 0},
				{0, 0, 0, 1, 0, 0, 0},
			}
		})
	})
}
