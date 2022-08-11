package crackingthecodinginterview_test

import (
	ctci "golang-coding-problems/internal/crackingthecodinginterview"

	"testing"

	"github.com/franela/goblin"
)

func TestGraph(t *testing.T) {
	g := goblin.Goblin(t)

	g.Describe("Graph data struct", func() {
		var graph *ctci.Graph

		g.It("Should find path to a connected node via DFS", func() {
			g.Assert(graph.HasPathDFS(0, 5)).IsTrue()
		})

		g.It("Should not find path to a not connected node via DFS", func() {
			g.Assert(graph.HasPathDFS(0, 6)).IsFalse()
		})

		g.It("Should find path to a connected node via BFS", func() {
			g.Assert(graph.HasPathBFS(0, 5)).IsTrue()
		})

		g.It("Should not find path to a not connected node via BFS", func() {
			g.Assert(graph.HasPathBFS(0, 6)).IsFalse()
		})

		g.Before(func() {
			graph = ctci.MakeGraph()

			for i := 0; i <= 6; i++ {
				graph.AddNode(&ctci.GraphNode{
					Value: i,
				})
			}

			graph.AddEdge(0, 1)
			graph.AddEdge(0, 2)
			graph.AddEdge(1, 4)
			graph.AddEdge(2, 3)
			graph.AddEdge(3, 5)

		})
	})
}
