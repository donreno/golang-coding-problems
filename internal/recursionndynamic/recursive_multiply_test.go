package recursionndynamic_test

import (
	r "golang-coding-problems/internal/recursionndynamic"

	"testing"

	goblin "github.com/franela/goblin"
)

func TestRecursiveMultiply(t *testing.T) {
	g := goblin.Goblin(t)

	g.Describe("Recursive multiply", func() {

		g.It("Should multiply numbers as expected", func() {
			g.Assert(r.RecursiveMultiply(3, 5)).Eql(15)
			g.Assert(r.RecursiveMultiply(7, 3)).Eql(21)
			g.Assert(r.RecursiveMultiply(10, 9)).Eql(90)
			g.Assert(r.RecursiveMultiply(31500, 0)).Eql(0)
			g.Assert(r.RecursiveMultiply(7, -3)).Eql(-21)
			g.Assert(r.RecursiveMultiply(-7, 3)).Eql(-21)
			g.Assert(r.RecursiveMultiply(-7, -3)).Eql(21)
		})
	})
}
