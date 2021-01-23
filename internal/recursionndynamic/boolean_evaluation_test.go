package recursionndynamic_test

import (
	r "golang-coding-problems/internal/recursionndynamic"
	"testing"

	goblin "github.com/franela/goblin"
)

func TestBooleanEvaluation(t *testing.T) {
	g := goblin.Goblin(t)

	g.Describe("Boolean evaluation", func() {

		g.It("Should count expected number of ways parenthesizing to make expression be true or false", func() {
			g.Assert(r.BooleanEvaluation("1^0|0|1", false)).Eql(2)
			g.Assert(r.BooleanEvaluation("0&0&0&1^1|0", true)).Eql(10)
			g.Assert(r.BooleanEvaluation("", true)).Eql(0)
		})
	})
}
