package graphsntrees_test

import (
	s "golang-coding-problems/internal/structs"
	"testing"

	gnt "golang-coding-problems/internal/graphsntrees"

	goblin "github.com/franela/goblin"
)

func TestRouteBetweenNodes(t *testing.T) {
	g := goblin.Goblin(t)

	g.Describe("Route Between Graph nodes", func() {
		var root *s.GraphNode

		g.BeforeEach(func() {
			root = s.MakeGraphNode(1, 1)

			n2, n3, n4, n5, n6 := s.MakeGraphNode(2, 2), s.MakeGraphNode(3, 3), s.MakeGraphNode(4, 4), s.MakeGraphNode(5, 5), s.MakeGraphNode(6, 6)

			root.AddNode(n2)
			root.AddNode(n3)

			n2.AddNode(n6)
			n3.AddNode(n4)
			n4.AddNode(n5)
		})

		g.It("Should return true if there is a route to a node", func() {

			g.Assert(gnt.RouteBetweenNodes(5, root)).IsTrue()
			g.Assert(gnt.RouteBetweenNodes(6, root)).IsTrue()
		})

		g.It("Should return false if there is no route to a node", func() {

			g.Assert(gnt.RouteBetweenNodes(7, root)).IsFalse()
		})

	})
}
