package linkedlist_test

import (
	l "golang-coding-problems/internal/linkedlist"
	s "golang-coding-problems/internal/structs"
	"testing"

	goblin "github.com/franela/goblin"
)

func TestPartition(t *testing.T) {
	g := goblin.Goblin(t)

	g.Describe("Linked list partition problem", func() {
		var ll *s.LinkedList

		g.BeforeEach(func() {
			ll = new(s.LinkedList)
		})

		g.It("Should partition linkedlist correctly", func() {
			ll.Add(3).Add(5).Add(8).Add(5).Add(10).Add(2).Add(1)

			lPartitioned := l.Partition(ll, 5)

			g.Assert(lPartitioned.Head.Val).Eql(1)
			g.Assert(lPartitioned.Head.Next.Val).Eql(2)
			g.Assert(lPartitioned.Head.Next.Next.Val).Eql(3)
			g.Assert(lPartitioned.Head.Next.Next.Next.Val).Eql(5)
			g.Assert(lPartitioned.Head.Next.Next.Next.Next.Val).Eql(8)
			g.Assert(lPartitioned.Head.Next.Next.Next.Next.Next.Val).Eql(5)
			g.Assert(lPartitioned.Head.Next.Next.Next.Next.Next.Next.Val).Eql(10)
		})

		g.It("Should not fail on empty list input", func() {
			lPartitioned := l.Partition(ll, 4)

			g.Assert(lPartitioned.Head).IsZero()
		})

	})
}
