package structs_test

import (
	s "golang-coding-problems/internal/structs"
	"testing"

	goblin "github.com/franela/goblin"
)

func TestStacks(t *testing.T) {
	g := goblin.Goblin(t)

	g.Describe("Stacks", func() {
		var st s.Stack

		g.BeforeEach(func() {
			st = s.MakeStack()
		})

		g.It("Should Top nil if Stack is new (empty)", func() {
			g.Assert(st.Top()).IsNil()
		})

		g.It("Should Pop nil if Stack is new (empty)", func() {
			g.Assert(st.Pop()).IsNil()
		})

		g.It("Should return true for Empty if Stack is new", func() {
			g.Assert(st.Empty()).IsTrue()
		})

		g.It("Should Push elements and then top and pop them properly", func() {
			st.Push(15)
			st.Push(21)
			st.Push(0)

			g.Assert(st.Top().(int)).Equal(0)
			g.Assert(st.Pop().(int)).Equal(0)
			g.Assert(st.Pop().(int)).Equal(21)
			g.Assert(st.Pop().(int)).Equal(15)
			g.Assert(st.Empty()).IsTrue()
		})
	})
}

func BenchmarkStack(b *testing.B) {
	st := s.MakeStack()

	for n := 0; n < b.N; n++ {
		st.Push("HELLO")
		st.Pop()
	}
}
