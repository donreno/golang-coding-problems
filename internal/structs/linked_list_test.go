package structs_test

import (
	"testing"

	goblin "github.com/franela/goblin"

	s "golang-coding-problems/internal/structs"
)

func TestLinkedList(t *testing.T) {
	g := goblin.Goblin(t)

	g.Describe("Linked List", func() {
		var ll *s.LinkedList

		g.BeforeEach(func() {
			ll = new(s.LinkedList)
		})

		g.It("Should create an empty linkedlist first", func() {
			g.Assert(ll.Empty()).IsTrue()
		})

		g.It("Should not be empty after adding one element", func() {
			ll.Add("something")

			g.Assert(ll.Empty()).IsFalse()
			g.Assert(ll.Size()).Eql(1)
			g.Assert(ll.Head()).Eql(ll.Tail())
		})

		g.It("Should change size after remove", func() {
			ll.Add("good")
			ll.Add("ok")
			ll.Add("bad")

			ll.Remove(2)

			g.Assert(ll.Size()).Eql(2)
		})

		g.It("Should be able to convert a LinkedList into a slice", func() {
			ll.Add(1)
			ll.Add(2)
			ll.Add(3)

			g.Assert(ll.ToSlice()).Eql([]interface{}{1, 2, 3})
		})
	})
}