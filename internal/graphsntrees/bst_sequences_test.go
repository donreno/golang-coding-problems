package graphsntrees_test

import (
	"testing"

	gnt "golang-coding-problems/internal/graphsntrees"

	goblin "github.com/franela/goblin"
)

func TestBSTSequences(t *testing.T) {
	g := goblin.Goblin(t)

	g.Describe("BST Sequences", func() {
		var root *gnt.BinaryTreeNode

		g.Before(func() {
			root = &gnt.BinaryTreeNode{
				Val: 2,
				Left: &gnt.BinaryTreeNode{
					Val: 1,
				},
				Right: &gnt.BinaryTreeNode{
					Val: 3,
				},
			}
		})

		g.It("Should return the correct arrays", func() {
			combinations := gnt.FindAllPossibleBSTSequences(root)

			g.Assert(len(combinations)).Eql(2)
			g.Assert(combinations[0]).Eql([]int{2, 1, 3})
			g.Assert(combinations[1]).Eql([]int{2, 3, 1})
		})
	})
}
