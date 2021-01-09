package stacks_test

import (
	s "golang-coding-problems/internal/stacks"
	"testing"

	goblin "github.com/franela/goblin"
)

func TestQueueViaStacks(t *testing.T) {
	g := goblin.Goblin(t)

	g.Describe("Queue via stacks", func() {
		var q *s.QueueViaStack

		g.BeforeEach(func() {
			q = s.MakeQueueViaStack()
		})

		g.It("Should behave as a regular queue", func() {
			q.Enqueue(1)
			q.Enqueue(2)
			q.Enqueue(3)

			g.Assert(q.Head()).Eql(1)
			g.Assert(q.Dequeue()).Eql(1)
			g.Assert(q.Dequeue()).Eql(2)
			g.Assert(q.Dequeue()).Eql(3)

		})
	})
}
