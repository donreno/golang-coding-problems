package sorting_test

import (
	s "golang-coding-problems/internal/sorting"
	"testing"

	goblin "github.com/franela/goblin"
)

func TestRankFromStream(t *testing.T) {
	g := goblin.Goblin(t)

	g.Describe("Rank From Stream", func() {
		var stream *s.BinaryTreeNode

		g.It("Should return the correct rank for given number", func() {
			g.Assert(s.RankFromStream(stream, 1)).Eql(0)
			g.Assert(s.RankFromStream(stream, 3)).Eql(1)
			g.Assert(s.RankFromStream(stream, 4)).Eql(2)
			g.Assert(s.RankFromStream(stream, 5)).Eql(4)
		})

		g.BeforeEach(func() {
			stream = s.MakeBinaryTree()
			stream.Insert(5)
			stream.Insert(1)
			stream.Insert(4)
			stream.Insert(4)
			stream.Insert(5)
			stream.Insert(9)
			stream.Insert(7)
			stream.Insert(13)
			stream.Insert(3)
		})
	})
}
