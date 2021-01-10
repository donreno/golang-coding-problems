package graphsntrees_test

import (
	"testing"

	gnt "golang-coding-problems/internal/graphsntrees"

	goblin "github.com/franela/goblin"
)

func TestListOfDepths(t *testing.T) {
	g := goblin.Goblin(t)

	g.Describe("List of Depths", func() {

		root := &gnt.BinaryTreeNode{
			Val: 4,
			Left: &gnt.BinaryTreeNode{
				Val:   2,
				Left:  &gnt.BinaryTreeNode{Val: 1},
				Right: &gnt.BinaryTreeNode{Val: 3},
			},
			Right: &gnt.BinaryTreeNode{
				Val:   6,
				Left:  &gnt.BinaryTreeNode{Val: 5},
				Right: &gnt.BinaryTreeNode{Val: 7},
			},
		}

		g.It("Should Create expected list of depth", func() {
			lod := gnt.ListOfDepths(root)

			g.Assert(lod[0].Val).Eql(4)
			g.Assert(lod[1].Val).Eql(6)
			g.Assert(lod[1].Next.Val).Eql(2)
			g.Assert(lod[2].Val).Eql(7)
			g.Assert(lod[2].Next.Val).Eql(5)
			g.Assert(lod[2].Next.Next.Val).Eql(3)
			g.Assert(lod[2].Next.Next.Next.Val).Eql(1)
		})

	})
}
