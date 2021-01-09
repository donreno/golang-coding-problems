package stacks_test

import (
	s "golang-coding-problems/internal/stacks"
	"testing"

	goblin "github.com/franela/goblin"
)

func TestMinStack(t *testing.T) {
	g := goblin.Goblin(t)

	g.Describe("Min Stack", func() {
		var minStack *s.MinStack

		g.BeforeEach(func() {
			minStack = new(s.MinStack)
		})

		g.It("Should show correct Top element when Peeking and Popping", func() {
			minStack.Push(1)
			minStack.Push(5)

			g.Assert(minStack.Peek()).Eql(5)
			g.Assert(minStack.Pop()).Eql(5)
			g.Assert(minStack.Peek()).Eql(1)
		})

		g.It("Should return correct Min element", func() {
			minStack.Push(2)
			minStack.Push(-2)
			minStack.Push(1)
			minStack.Push(-5)

			g.Assert(minStack.Min()).Eql(-5)
			g.Assert(minStack.Pop()).Eql(-5)
			g.Assert(minStack.Min()).Eql(-2)

		})

	})
}
