package graphsntrees_test

import (
	"testing"

	gnt "golang-coding-problems/internal/graphsntrees"

	goblin "github.com/franela/goblin"
)

func TestBuildOrder(t *testing.T) {
	g := goblin.Goblin(t)

	g.Describe("Build Order", func() {
		var depGraph map[rune][]rune

		g.Before(func() {
			depGraph = map[rune][]rune{
				'd': []rune{'a', 'b'},
				'b': []rune{'f'},
				'a': []rune{'f'},
				'c': []rune{'d'},
			}
		})

		g.It("Should return correct build order", func() {
			buildOrder := gnt.BuildOrder([]rune{'a', 'b', 'c', 'd', 'e', 'f'}, depGraph)

			g.Assert(buildOrder[0]).Eql('f')
			g.Assert(buildOrder[1]).Eql('a')
			g.Assert(buildOrder[2]).Eql('b')
			g.Assert(buildOrder[3]).Eql('d')
			g.Assert(buildOrder[4]).Eql('c')
			g.Assert(buildOrder[5]).Eql('e')
		})
	})
}
