package linkedlist_test

import (
	l "golang-coding-problems/internal/linkedlist"
	s "golang-coding-problems/internal/structs"
	"testing"

	goblin "github.com/franela/goblin"
)

func TestRemoveDups(t *testing.T) {
	g := goblin.Goblin(t)

	g.Describe("Remove Dups from Linkedlist", func() {
		var ll *s.LinkedList

		g.BeforeEach(func() {
			ll = new(s.LinkedList)
		})

		g.It("Should remove all the dups from the linkedlist", func() {
			ll.Add(1).Add(2).Add(3).Add(1).Add(3).Add(4)

			l.RemoveDups(ll)

			llSlice := ll.ToSlice()

			g.Assert(llSlice[0]).Eql(1)
			g.Assert(llSlice[1]).Eql(2)
			g.Assert(llSlice[2]).Eql(3)
			g.Assert(llSlice[3]).Eql(4)
		})

		g.It("Should not remove any elements if no duplicates", func() {
			ll.Add(1).Add(2).Add(3)

			l.RemoveDups(ll)

			llSlice := ll.ToSlice()

			g.Assert(llSlice[0]).Eql(1)
			g.Assert(llSlice[1]).Eql(2)
			g.Assert(llSlice[2]).Eql(3)
		})
	})
}
