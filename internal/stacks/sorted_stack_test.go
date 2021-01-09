package stacks_test

import (
	s "golang-coding-problems/internal/stacks"
	"testing"

	goblin "github.com/franela/goblin"
)

func TestSortedStacks(t *testing.T) {
	g := goblin.Goblin(t)

	g.Describe("Sorted stack", func() {
		var sortedStack *s.SortedStack

		g.BeforeEach(func() {
			sortedStack = s.MakeSortedStack()
		})

		g.It("Should sort stack every time an element is pushed", func() {

			sortedStack.Push(5)

			g.Assert(sortedStack.Peek()).Eql(5)

			sortedStack.Push(3)

			g.Assert(sortedStack.Peek()).Eql(3)

			sortedStack.Push(7)

			g.Assert(sortedStack.Peek()).Eql(3)

			g.Assert(sortedStack.Pop()).Eql(3)
			g.Assert(sortedStack.Pop()).Eql(5)
			g.Assert(sortedStack.Pop()).Eql(7)
		})
	})
}
