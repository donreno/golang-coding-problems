package recursionndynamic_test

import (
	r "golang-coding-problems/internal/recursionndynamic"
	s "golang-coding-problems/internal/structs"

	"testing"

	goblin "github.com/franela/goblin"
)

func TestTowersOfHanoi(t *testing.T) {
	g := goblin.Goblin(t)

	g.Describe("Towers Of Hanoi", func() {
		var a, b, c s.Stack

		g.BeforeEach(func() {
			a, b, c = s.MakeStack(), s.MakeStack(), s.MakeStack()
		})

		g.It("Should solve base case properly", func() {
			a.Push(1)

			r.TowersOfHanoi(1, a, c, b)

			g.Assert(c.Pop().(int)).Eql(1)
			g.Assert(a.Empty()).IsTrue()
			g.Assert(b.Empty()).IsTrue()
		})

		g.It("Should solve 2 elements case properly", func() {
			a.Push(2)
			a.Push(1)

			r.TowersOfHanoi(2, a, c, b)

			g.Assert(c.Pop().(int)).Eql(1)
			g.Assert(c.Pop().(int)).Eql(2)
			g.Assert(a.Empty()).IsTrue()
			g.Assert(b.Empty()).IsTrue()
		})

		g.It("Should solve 10 elements case properly", func() {
			a.Push(10)
			a.Push(9)
			a.Push(8)
			a.Push(7)
			a.Push(6)
			a.Push(5)
			a.Push(4)
			a.Push(3)
			a.Push(2)
			a.Push(1)

			r.TowersOfHanoi(10, a, c, b)

			g.Assert(c.Pop().(int)).Eql(1)
			g.Assert(c.Pop().(int)).Eql(2)
			g.Assert(c.Pop().(int)).Eql(3)
			g.Assert(c.Pop().(int)).Eql(4)
			g.Assert(c.Pop().(int)).Eql(5)
			g.Assert(c.Pop().(int)).Eql(6)
			g.Assert(c.Pop().(int)).Eql(7)
			g.Assert(c.Pop().(int)).Eql(8)
			g.Assert(c.Pop().(int)).Eql(9)
			g.Assert(c.Pop().(int)).Eql(10)
			g.Assert(a.Empty()).IsTrue()
			g.Assert(b.Empty()).IsTrue()
		})
	})
}
