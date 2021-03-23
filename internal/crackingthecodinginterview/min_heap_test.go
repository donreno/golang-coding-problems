package crackingthecodinginterview_test

import (
	ctci "golang-coding-problems/internal/crackingthecodinginterview"

	"testing"

	"github.com/franela/goblin"
)

func TestMinHeap(t *testing.T) {
	g := goblin.Goblin(t)

	g.Describe("Min Heap", func() {
		var heap *ctci.MinHeap

		g.It("Should operate min heap correctly and poll expected values on specific order", func() {

			heap.Add(20)
			heap.Add(90)
			heap.Add(2)
			heap.Add(9)
			heap.Add(1)
			heap.Add(50)

			g.Assert(heap.Poll()).Equal(1)
			g.Assert(heap.Poll()).Equal(2)
			g.Assert(heap.Poll()).Equal(9)
			g.Assert(heap.Poll()).Equal(20)
			g.Assert(heap.Poll()).Equal(50)
			g.Assert(heap.Poll()).Equal(90)
		})

		g.BeforeEach(func() {
			heap = ctci.BuildMinHeap(3)
		})
	})
}
