package golang_coding_problems_test

import (
	c "golang-coding-problems"
	"testing"

	goblin "github.com/franela/goblin"
)

func TestGraph(t *testing.T) {
	g := goblin.Goblin(t)

	g.Describe("Graphs", func() {
		var graph *c.GraphNode

		g.BeforeEach(func() {
			graph = c.MakeGraphNode(1, 1)
		})

		g.It("Should add Nodes properly", func() {
			otherNode := c.MakeGraphNode(2, 2)

			graph.AddNode(otherNode)

			g.Assert(len(graph.Nodes)).IsNotZero()
		})

		g.It("Should get a node if added", func() {
			otherNode := c.MakeGraphNode(2, 2)

			graph.AddNode(otherNode)

			foundNode := graph.GetNode(2)

			g.Assert(foundNode).Equal(otherNode)
		})

		g.It("Should not get a node if not added", func() {
			foundNode := graph.GetNode(5)

			g.Assert(foundNode).IsZero()
		})
	})

}

func BenchmarkGraph(b *testing.B) {
	g := c.MakeGraphNode(1, 1)
	node1 := c.MakeGraphNode(2, 2)
	node2 := c.MakeGraphNode(3, 3)
	node3 := c.MakeGraphNode(4, 4)

	node1.AddNode(node3)
	node3.AddNode(node2)
	node2.AddNode(g)

	g.AddNode(node1)

	for n := 0; n < b.N; n++ {
		g.GetNode(3)
	}
}
