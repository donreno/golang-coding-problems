package googleinterview_test

import (
	gi "golang-coding-problems/internal/googleinterview"

	"testing"

	"github.com/franela/goblin"
)

func TestSwapLeftAndRightTreeChild(t *testing.T) {

	g := goblin.Goblin(t)

	g.Describe("Swap Left and Right Tree childs", func() {
		tree := gi.MakeTreeNode(4)
		tree.Left = gi.MakeTreeNode(2)
		tree.Left.Left = gi.MakeTreeNode(1)
		tree.Left.Right = gi.MakeTreeNode(3)
		tree.Right = gi.MakeTreeNode(6)
		tree.Right.Left = gi.MakeTreeNode(5)
		tree.Right.Right = gi.MakeTreeNode(7)

		g.It("Should swap left and right childs from tree", func() {
			gi.SwapTreeChildNodes(tree)

			g.Assert(tree.Value).Eql(4)
			g.Assert(tree.Left.Value).Eql(6)
			g.Assert(tree.Left.Left.Value).Eql(7)
			g.Assert(tree.Left.Right.Value).Eql(5)
			g.Assert(tree.Right.Value).Eql(2)
			g.Assert(tree.Right.Left.Value).Eql(3)
			g.Assert(tree.Right.Right.Value).Eql(1)
		})
	})
}
