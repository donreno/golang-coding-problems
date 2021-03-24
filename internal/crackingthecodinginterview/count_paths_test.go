package crackingthecodinginterview_test

import (
	ctci "golang-coding-problems/internal/crackingthecodinginterview"

	"testing"

	"github.com/franela/goblin"
)

func TestCountPathsForSteps(t *testing.T) {
	g := goblin.Goblin(t)

	g.Describe("Count Paths", func() {
		g.It("Should return expected number of paths to reach top of the stair", func() {
			g.Assert(ctci.CountPaths(0)).Eql(0)
			g.Assert(ctci.CountPaths(1)).Eql(1)
			g.Assert(ctci.CountPaths(2)).Eql(2)
			g.Assert(ctci.CountPaths(3)).Eql(3)
			g.Assert(ctci.CountPaths(4)).Eql(6)
			g.Assert(ctci.CountPaths(5)).Eql(11)
		})
	})
}
