package crackingthecodinginterview_test

import (
	ctci "golang-coding-problems/internal/crackingthecodinginterview"

	"testing"

	"github.com/franela/goblin"
)

func TestCountPaths(t *testing.T) {
	g := goblin.Goblin(t)

	g.Describe("Count Paths to exit a grid using dynamic programming", func() {
		var grid [][]int

		g.It("Should find the right number of ways to get to the exit", func() {
			g.Assert(ctci.FindPathsDynamic(grid)).Eql(27)
		})

		g.Before(func() {
			grid = [][]int{
				{0, 0, 0, 0, 0, 0, 0, 0},
				{0, 0, -1, 0, 0, 0, -1, 0},
				{0, 0, 0, 0, -1, 0, 0, 0},
				{-1, 0, -1, 0, 0, -1, 0, 0},
				{0, 0, -1, 0, 0, 0, 0, 0},
				{0, 0, 0, -1, -1, 0, -1, 0},
				{0, -1, 0, 0, 0, -1, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 0},
			}
		})
	})
}
