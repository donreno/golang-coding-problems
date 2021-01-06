package linkedlist_test

import (
	l "golang-coding-problems/internal/linkedlist"
	s "golang-coding-problems/internal/structs"
	"testing"

	goblin "github.com/franela/goblin"
)

func TestLoopDetection(t *testing.T) {
	g := goblin.Goblin(t)

	g.Describe("Loop detection on Linkedlist", func() {

		var ll *s.LinkedList

		g.BeforeEach(func() {
			ll = new(s.LinkedList)
		})

		g.It("Should return true if there is a loop on linkedlist", func() {

			nodeA := s.MakeListNode('A')
			nodeB := s.MakeListNode('B')
			nodeC := s.MakeListNode('C')
			nodeD := s.MakeListNode('D')
			nodeE := s.MakeListNode('E')

			nodeA.Next = nodeB
			nodeB.Next = nodeC
			nodeC.Next = nodeD
			nodeD.Next = nodeE
			nodeE.Next = nodeC

			ll.Head = nodeA

			g.Assert(l.HasLoop(ll)).IsTrue()
		})

		g.It("Should return false if there is no loop", func() {
			nodeA := s.MakeListNode('A')
			nodeB := s.MakeListNode('B')
			nodeC := s.MakeListNode('C')

			nodeA.Next = nodeB
			nodeB.Next = nodeC

			ll.Head = nodeA

			g.Assert(l.HasLoop(ll)).IsFalse()
		})

		g.It("Should return false if Linkedlist is empty", func() {
			g.Assert(l.HasLoop(ll)).IsFalse()
		})
	})
}
