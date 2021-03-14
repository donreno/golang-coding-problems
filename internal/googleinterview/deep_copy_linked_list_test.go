package googleinterview_test

import (
	gi "golang-coding-problems/internal/googleinterview"

	"testing"

	"github.com/franela/goblin"
)

func TestDeepCopyLinkedList(t *testing.T) {
	g := goblin.Goblin(t)

	g.Describe("Deep copy linked list", func() {
		linkedList := new(gi.SpecialLL)
		g.It("Should deep copy linkedlist", func() {
			linkedList.Data = 1
			linkedList.Next = new(gi.SpecialLL)
			linkedList.Next.Data = 2
			linkedList.Next.Arbitrary = linkedList

			llCopy := gi.DeepCopyLinkedList(linkedList)

			g.Assert(llCopy.Data).Eql(1)
			g.Assert(llCopy.Next.Data).Eql(2)
			g.Assert(llCopy.Next.Arbitrary.Data).Eql(1)
		})
	})
}
