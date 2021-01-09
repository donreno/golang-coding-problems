package stacks_test

import (
	s "golang-coding-problems/internal/stacks"
	"testing"

	goblin "github.com/franela/goblin"
)

func TestStackOfPlates(t *testing.T) {
	g := goblin.Goblin(t)

	g.Describe("Stack of Plates", func() {
		var stackOfPlates *s.StackOfPlates

		g.BeforeEach(func() {
			stackOfPlates = s.MakeStackOfPlates(3)
		})

		g.It("Stack of plates should behave as description says", func() {
			stackOfPlates.Push(1)
			stackOfPlates.Push(2)
			stackOfPlates.Push(3)
			stackOfPlates.Push(4)
			stackOfPlates.Push(5)
			stackOfPlates.Push(6)
			stackOfPlates.Push(7)
			stackOfPlates.Push(8)

			g.Assert(stackOfPlates.Pop()).Eql(8)
			g.Assert(stackOfPlates.Pop()).Eql(7)
			g.Assert(stackOfPlates.Pop()).Eql(6)
			g.Assert(stackOfPlates.Pop()).Eql(5)
			g.Assert(stackOfPlates.Pop()).Eql(4)
			g.Assert(stackOfPlates.Pop()).Eql(3)
			g.Assert(stackOfPlates.Pop()).Eql(2)
			g.Assert(stackOfPlates.Peek()).Eql(1)
			g.Assert(stackOfPlates.Pop()).Eql(1)
			g.Assert(stackOfPlates.Peek()).Eql(-1)
			g.Assert(stackOfPlates.Pop()).Eql(-1)
		})
	})
}
