package crackingthecodinginterview

// Graph Graph data struct
type Graph struct {
	nodeLookup map[int]*GraphNode
}

// GraphNode a Graph node
type GraphNode struct {
	ID     int
	Value  int
	Childs []*GraphNode
}

// MakeGraph makes a new Graph
func MakeGraph() *Graph {
	return &Graph{
		nodeLookup: make(map[int]*GraphNode),
	}
}

// GetNode gets a node by id
func (g *Graph) GetNode(id int) *GraphNode {
	return g.nodeLookup[id]
}

// AddNode adds node to graph
func (g *Graph) AddNode(node *GraphNode) {
	if node == nil {
		return
	}

	g.nodeLookup[node.ID] = node
}

// AddEdge adds an edge to a given source node to link to destination node
func (g *Graph) AddEdge(source, destination int) {
	srcNode := g.nodeLookup[source]
	destNode := g.nodeLookup[destination]

	if srcNode != nil && destNode != nil {
		srcNode.Childs = append(srcNode.Childs, destNode)
	}
}

// HasPathDFS checks if there is a path to a node with DFS algorithm
func (g *Graph) HasPathDFS(source, destination int) bool {
	srcNode := g.nodeLookup[source]
	destNode := g.nodeLookup[destination]
	return g.hasPathDFS(srcNode, destNode, make(map[int]bool))
}

func (g *Graph) hasPathDFS(source, destination *GraphNode, visited map[int]bool) bool {
	if visited[source.ID] {
		return false
	}

	if source.ID == destination.ID {
		return true
	}

	visited[source.ID] = true

	for _, child := range source.Childs {
		if g.hasPathDFS(child, destination, visited) {
			return true
		}
	}

	return false
}

// HasPathBFS has path to a node using BFS
func (g *Graph) HasPathBFS(source, destination int) bool {
	srcNode := g.nodeLookup[source]
	destNode := g.nodeLookup[destination]
	nextToVisit := []*GraphNode{srcNode}
	visited := make(map[int]bool)

	var current *GraphNode

	for len(nextToVisit) > 0 {
		current, nextToVisit = nextToVisit[0], nextToVisit[1:]

		if current.ID == destNode.ID {
			return true
		}

		if visited[current.ID] {
			continue
		}

		visited[current.ID] = true

		for _, child := range current.Childs {
			nextToVisit = append(nextToVisit, child)
		}
	}

	return false
}
