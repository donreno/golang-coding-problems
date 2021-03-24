package linkedlist_test

import (
	l "golang-coding-problems/internal/linkedlist"
	s "golang-coding-problems/internal/structs"
	"time"

	"testing"

	goblin "github.com/franela/goblin"
)

func TestMergeKSortedLinkedLists(t *testing.T) {
	g := goblin.Goblin(t)

	g.Describe("Merge K Sorted LinkedLists", func() {
		var kLists []*s.ListNode

		g.It("Merge K sorted LL should result in expected merged list", func() {
			g.Timeout(5 * time.Minute)
			mergedLL := l.MergeKSortedLinkedLists(kLists)

			g.Assert(mergedLL.Val).Eql(1)
			g.Assert(mergedLL.Next.Val).Eql(2)
			g.Assert(mergedLL.Next.Next.Val).Eql(3)
			g.Assert(mergedLL.Next.Next.Next.Val).Eql(4)
			g.Assert(mergedLL.Next.Next.Next.Next.Val).Eql(5)
			g.Assert(mergedLL.Next.Next.Next.Next.Next.Val).Eql(6)
		})

		g.BeforeEach(func() {
			kLists = make([]*s.ListNode, 3)
			kLists[0], kLists[1], kLists[2] = &s.ListNode{Val: 2}, &s.ListNode{Val: 1}, &s.ListNode{Val: 5}

			kLists[0].Next, kLists[1].Next, kLists[2].Next = &s.ListNode{Val: 3}, &s.ListNode{Val: 4}, &s.ListNode{Val: 6}
		})

	})
}
