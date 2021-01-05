package linkedlist_test

import (
	l "golang-coding-problems/internal/linkedlist"
	s "golang-coding-problems/internal/structs"
	"testing"

	goblin "github.com/franela/goblin"
)

func TestSumLists(t *testing.T) {
	g := goblin.Goblin(t)

	g.Describe("Sum linkedlists", func() {

		g.It("Should sum two numbers same digits correctly", func() {
			l1 := new(s.LinkedList).Add(7).Add(1).Add(6)
			l2 := new(s.LinkedList).Add(5).Add(9).Add(2)

			sum := l.SumLists(l1, l2)

			g.Assert(sum.Get(0)).Eql(2)
			g.Assert(sum.Get(1)).Eql(1)
			g.Assert(sum.Get(2)).Eql(9)
		})

		g.It("Should return l2 if l1 is empty", func() {
			l1 := new(s.LinkedList)
			l2 := new(s.LinkedList).Add(5).Add(9).Add(2)

			sum := l.SumLists(l1, l2)

			g.Assert(sum.Get(0)).Eql(5)
			g.Assert(sum.Get(1)).Eql(9)
			g.Assert(sum.Get(2)).Eql(2)
		})

		g.It("Should return l1 if l2 is empty", func() {
			l1 := new(s.LinkedList).Add(7).Add(1).Add(6)
			l2 := new(s.LinkedList)

			sum := l.SumLists(l1, l2)

			g.Assert(sum.Get(0)).Eql(7)
			g.Assert(sum.Get(1)).Eql(1)
			g.Assert(sum.Get(2)).Eql(6)
		})
	})

}
