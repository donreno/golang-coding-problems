package structs_test

import (
	s "golang-coding-problems/internal/structs"
	"testing"

	goblin "github.com/franela/goblin"
)

func TestQueue(t *testing.T) {
	g := goblin.Goblin(t)

	g.Describe("Queues", func() {
		var q s.Queue

		g.BeforeEach(func() {
			q = s.MakeQueue()
		})

		g.It("Should create empty queue", func() {
			g.Assert(q.Empty()).IsTrue()
		})

		g.It("Should have first enqueued element as Head and should not be empty", func() {
			q.Enqueue(1)
			q.Enqueue(2)

			g.Assert(q.Head()).Eql(1)
			g.Assert(q.Empty()).IsFalse()
		})

		g.It("Should dequeue in proper order", func() {
			q.Enqueue("A")
			q.Enqueue("B")
			q.Enqueue("C")

			g.Assert(q.Dequeue()).Eql("A")
			g.Assert(q.Dequeue()).Eql("B")
			g.Assert(q.Head()).Eql("C")
		})
	})
}

func BenchmarkQueue(b *testing.B) {
	q := s.MakeQueue()

	for n := 0; n < b.N; n++ {
		q.Enqueue("HELLO")
		q.Dequeue()
	}
}
