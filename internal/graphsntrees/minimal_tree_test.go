package graphsntrees_test

import (
	"testing"

	gnt "golang-coding-problems/internal/graphsntrees"

	goblin "github.com/franela/goblin"
)

func TestMinimalTree(t *testing.T) {
	g := goblin.Goblin(t)

	g.Describe("Minimal Tree", func() {

		g.It("Should create minimal tree according to criteria (Test 1)", func() {
			elems := []int{1, 2, 3, 4, 5, 6, 7}

			root := gnt.FillMinimalBinaryTree(elems)

			g.Assert(root.Val).Eql(4)
			g.Assert(root.Left.Val).Eql(2)
			g.Assert(root.Left.Left.Val).Eql(1)
			g.Assert(root.Left.Right.Val).Eql(3)
			g.Assert(root.Right.Val).Eql(6)
			g.Assert(root.Right.Left.Val).Eql(5)
			g.Assert(root.Right.Right.Val).Eql(7)
		})

		g.It("Should create minimal tree according to criteria (Test 2)", func() {
			elems := []int{1, 2, 4, 5, 6, 7}

			root := gnt.FillMinimalBinaryTree(elems)

			g.Assert(root.Val).Eql(5)
			g.Assert(root.Left.Val).Eql(2)
			g.Assert(root.Left.Left.Val).Eql(1)
			g.Assert(root.Left.Right.Val).Eql(4)
			g.Assert(root.Right.Val).Eql(7)
			g.Assert(root.Right.Left.Val).Eql(6)
		})
	})
}
