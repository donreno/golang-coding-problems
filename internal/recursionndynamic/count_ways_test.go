package recursionndynamic_test

import (
	r "golang-coding-problems/internal/recursionndynamic"
	"testing"

	goblin "github.com/franela/goblin"
)

func TestCountWays(t *testing.T) {
	g := goblin.Goblin(t)

	g.Describe("Count Ways to get trough stairs", func() {

		g.It("It should return right number of ways for given stair steps", func() {
			g.Assert(r.CountWays(1)).Eql(1)
			g.Assert(r.CountWays(2)).Eql(2)
			g.Assert(r.CountWays(3)).Eql(4)
			g.Assert(r.CountWays(4)).Eql(7)
		})

	})
}
