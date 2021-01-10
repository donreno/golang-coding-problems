package graphsntrees

import (
	s "golang-coding-problems/internal/structs"
)

// RouteBetweenNodes find route between nodes BFS
func RouteBetweenNodes(routeToID int, node *s.GraphNode) bool {
	visited := make(map[int]bool)
	q := s.MakeQueue()
	visited[node.ID] = true

	q.Enqueue(node)

	for !q.Empty() {
		n := q.Dequeue().(*s.GraphNode)

		if n.ID == routeToID {
			return true
		}

		for _, a := range n.Nodes {

			visited[a.ID] = true
			q.Enqueue(a)
		}
	}

	return false
}
