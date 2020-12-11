package golang_coding_problems_test

import (
	c "golang-coding-problems"
	"testing"

	goblin "github.com/franela/goblin"
)

func TestStacks(t *testing.T) {
	g := goblin.Goblin(t)

	g.Describe("Stacks", func() {
		var s c.Stack

		g.BeforeEach(func() {
			s = c.MakeStack()
		})

		g.It("Should Top nil if Stack is new (empty)", func() {
			g.Assert(s.Top()).IsNil()
		})

		g.It("Should Pop nil if Stack is new (empty)", func() {
			g.Assert(s.Pop()).IsNil()
		})

		g.It("Should return true for Empty if Stack is new", func() {
			g.Assert(s.Empty()).IsTrue()
		})

		g.It("Should Push elements and then top and pop them properly", func() {
			s.Push(15)
			s.Push(21)
			s.Push(0)

			g.Assert(s.Top().(int)).Equal(0)
			g.Assert(s.Pop().(int)).Equal(0)
			g.Assert(s.Pop().(int)).Equal(21)
			g.Assert(s.Pop().(int)).Equal(15)
			g.Assert(s.Empty()).IsTrue()
		})
	})
}

func BenchmarkStack(b *testing.B) {
	s := c.MakeStack()

	for n := 0; n < b.N; n++ {
		s.Push("HELLO")
		s.Pop()
	}
}
