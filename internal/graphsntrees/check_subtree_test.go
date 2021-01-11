package graphsntrees_test

import (
	"testing"

	gnt "golang-coding-problems/internal/graphsntrees"

	goblin "github.com/franela/goblin"
)

func TestCheckSubtree(t *testing.T) {
	g := goblin.Goblin(t)

	g.Describe("Check Subtree", func() {
		var t1, t2 *gnt.BinaryTreeNode

		g.Before(func() {
			t1 = &gnt.BinaryTreeNode{
				Val: 4,
				Left: &gnt.BinaryTreeNode{
					Val:   2,
					Left:  &gnt.BinaryTreeNode{Val: 1},
					Right: &gnt.BinaryTreeNode{Val: 3},
				},
				Right: &gnt.BinaryTreeNode{
					Val:  6,
					Left: &gnt.BinaryTreeNode{Val: 5},
				},
			}

			t2 = &gnt.BinaryTreeNode{
				Val:   2,
				Left:  &gnt.BinaryTreeNode{Val: 1},
				Right: &gnt.BinaryTreeNode{Val: 3},
			}
		})

		g.It("Should identify if t2 is a subtree", func() {
			g.Assert(gnt.CheckT2IsSubtreeOfT1(t1, t2)).IsTrue()
		})

		g.It("Should return true if t2 is nil", func() {
			g.Assert(gnt.CheckT2IsSubtreeOfT1(t1, nil)).IsTrue()
		})

		g.It("Should return true if t2 =  t1", func() {
			g.Assert(gnt.CheckT2IsSubtreeOfT1(t1, t1)).IsTrue()
		})

		g.It("Should return false if t2's root is not even on t1", func() {
			t2 = &gnt.BinaryTreeNode{
				Val: -1000,
			}

			g.Assert(gnt.CheckT2IsSubtreeOfT1(t1, t2)).IsFalse()
		})

		g.It("Should return false if t2 is not subtree of t1 (e.g 1)", func() {
			t2 = &gnt.BinaryTreeNode{
				Val:  2,
				Left: &gnt.BinaryTreeNode{Val: 1},
			}

			g.Assert(gnt.CheckT2IsSubtreeOfT1(t1, t2)).IsFalse()
		})

		g.It("Should return false if t2 is not subtree of t1 (e.g 2)", func() {
			t2 = &gnt.BinaryTreeNode{
				Val:   2,
				Left:  &gnt.BinaryTreeNode{Val: 1},
				Right: &gnt.BinaryTreeNode{Val: 4},
			}

			g.Assert(gnt.CheckT2IsSubtreeOfT1(t1, t2)).IsFalse()
		})
	})
}
