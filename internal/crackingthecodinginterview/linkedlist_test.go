package crackingthecodinginterview_test

import (
	ctci "golang-coding-problems/internal/crackingthecodinginterview"

	"testing"

	"github.com/franela/goblin"
)

func TestLinkedList(t *testing.T) {
	g := goblin.Goblin(t)

	g.Describe("Linked list", func() {
		var ll *ctci.LinkedList

		g.It("Should append, prepend and delete elements properly", func() {
			ll.Append(2)
			ll.Append(3)
			ll.Prepend(1)

			g.Assert(ll.Head.Data).Eql(1)

			ll.DeleteWithValue(2)
			ll.DeleteWithValue(1)

			g.Assert(ll.Head.Data).Eql(3)
		})

		g.BeforeEach(func() {
			ll = new(ctci.LinkedList)
		})
	})
}
