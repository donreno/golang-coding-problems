package googleinterview_test

import (
	gi "golang-coding-problems/internal/googleinterview"

	"testing"

	"github.com/franela/goblin"
)

func TestCompareBinaryTrees(t *testing.T) {

	g := goblin.Goblin(t)

	g.Describe("Compare binary trees", func() {
		tree1 := gi.MakeTreeNode(4)
		tree1.Left = gi.MakeTreeNode(2)
		tree1.Left.Left = gi.MakeTreeNode(1)
		tree1.Left.Right = gi.MakeTreeNode(3)
		tree1.Right = gi.MakeTreeNode(6)
		tree1.Right.Left = gi.MakeTreeNode(5)
		tree1.Right.Right = gi.MakeTreeNode(7)

		tree2 := gi.MakeTreeNode(4)
		tree2.Left = gi.MakeTreeNode(2)
		tree2.Left.Left = gi.MakeTreeNode(1)
		tree2.Left.Right = gi.MakeTreeNode(3)
		tree2.Right = gi.MakeTreeNode(6)
		tree2.Right.Left = gi.MakeTreeNode(5)
		tree2.Right.Right = gi.MakeTreeNode(7)

		tree3 := gi.MakeTreeNode(4)
		tree3.Left = gi.MakeTreeNode(2)
		tree3.Left.Left = gi.MakeTreeNode(1)
		tree3.Left.Right = gi.MakeTreeNode(3)
		tree3.Right = gi.MakeTreeNode(6)
		tree3.Right.Left = gi.MakeTreeNode(5)
		tree3.Right.Right = gi.MakeTreeNode(9)

		g.It("Should return true if trees are equal", func() {
			g.Assert(gi.CompareTrees(tree1, tree2)).IsTrue()
		})

		g.It("Should return false if trees are not equal", func() {
			g.Assert(gi.CompareTrees(tree1, tree2)).IsTrue()
		})
	})
}
